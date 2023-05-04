import { CHECKLIST_PENCAIRAN_URL, RUPIAH } from "../../../const";
import { useStore } from "../../../Context";
import Swal from 'sweetalert2';
import { StatusCodes, getReasonPhrase } from 'http-status-codes';
import { useEffect, useRef, useState } from 'react';
import Table from 'react-bootstrap/Table';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Pagination, { bootstrap5PaginationPreset } from 'react-responsive-pagination';

export default function ChecklistPencairan() {
	const { dispatch } = useStore();
	const handleLogOut = () => dispatch({ type: 'logout' });
	const [records, setRecords] = useState([]);
	const [countRecord, setCountRecord] = useState(0);
	const [countPage, setCountPage] = useState(0);
	const [page, setPage] = useState(1);
	const limit = useRef(5);

	useEffect(() => getResource(page, limit.current), [page, limit.current]);

	const getResource = (page, limit) => {
		fetch(`${CHECKLIST_PENCAIRAN_URL}?page=${page}&limit=${limit}`, {
			method: 'GET',
			credentials: 'include',
			headers: {
				'Accept': 'application/json'
			}
		})
			.then(res => {
				if (!res.ok && (res.status === StatusCodes.UNAUTHORIZED || res.status === StatusCodes.FORBIDDEN)) {
					handleLogOut();
					throw new Error(`${res.status}::${res.statusText}`);
				}
				return res.json()
			})
			.then(json => {
				if (!json.data) {
					Swal.fire({ icon: 'error', title: 'Error', text: getReasonPhrase(json.code) });
					throw new Error(`${json.code}::${getReasonPhrase(json.code)}::${json.message}`);
				}
				if (json.data.records) {
					const newRecords = json.data.records.map((record) => ({ ...record, is_checked: false }));
					setRecords(newRecords);
				} else {
					setRecords([]);
				}
				setCountRecord(json.data.count_record);
				setCountPage(json.data.count_page);
			})
			.catch(err => console.error(err));
	};

	const handleChecked = (index) => {
		console.log(index);
		console.log(records[index].name);
		const newRecords = records.map((record, i) => {
			if (i === index) {
				record.is_checked = !record.is_checked
				return record
			}
			return record
		});
		console.log(newRecords[index].name);
		console.log(newRecords[index].is_checked);
		setRecords(newRecords);
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		const body = {
			custcodes: records.filter((record) => record.is_checked).map((record) => {
				return record.custcode
			})
		};
		if (body.custcodes.length === 0) {
			Swal.fire({ icon: 'error', title: 'Error', text: 'No Record(s) Submitted' });
			return
		}
		fetch(CHECKLIST_PENCAIRAN_URL, {
			method: 'PATCH',
			credentials: 'include',
			headers: {
				'Accept': 'application/json'
			},
			body: JSON.stringify(body)
		})
			.then(res => {
				if (!res.ok && (res.status === StatusCodes.UNAUTHORIZED || res.status === StatusCodes.FORBIDDEN)) {
					handleLogOut();
					throw new Error(`${res.status}::${res.statusText}`);
				}
				return res.json()
			})
			.then(json => {
				if (json.code !== StatusCodes.OK) {
					Swal.fire({ icon: 'error', title: 'Error', text: getReasonPhrase(json.code) });
					throw new Error(`${json.code}::${getReasonPhrase(json.code)}::${json.message}`);
				}
				Swal.fire({ icon: 'success', title: 'Success', showConfirmButton: false, timer: 1500 });
				if (page - 1 >= 1) {
					setPage(page - 1);
				} else {
					getResource(page, limit.current);
				}
			})
			.catch(err => console.error(err));
	};

	const recordsJSX = records.map((record, index) => {
		let no = (page - 1) * limit.current + index + 1;
		return (
			<tr key={index}>
				<td headers="no" style={{ fontSize: "1vw" }}>{no}</td>
				<td headers="ppk" style={{ fontSize: "1vw" }}>{record.ppk}</td>
				<td headers="name" style={{ fontSize: "1vw" }}>{record.name}</td>
				<td headers="channeling_company" style={{ fontSize: "1vw" }}>{record.channeling_company}</td>
				<td headers="drawdown_date" style={{ fontSize: "1vw" }}>{record.drawdown_date.substring(0, 10)}</td>
				<td headers="loan_amount" style={{ fontSize: "1vw" }}>{RUPIAH(record.loan_amount)}</td>
				<td headers="loan_period" style={{ fontSize: "1vw" }}>{record.loan_period}</td>
				<td headers="interest_effective" style={{ fontSize: "1vw" }}>{record.interest_effective}</td>
				<td headers="action" style={{ fontSize: "1vw" }} >
					<Form.Group controlId="formBasicCheckbox">
						<Form.Check type="checkbox" label="Pilih" name="custcodes[]" value={record.custcode} onChange={() => handleChecked(index)} checked={record.is_checked} />
					</Form.Group>
				</td>
			</tr>
		)
	});

	return (
		<>
			<div className="mt-3">
				<Pagination
					{...bootstrap5PaginationPreset}
					current={page}
					total={countPage}
					onPageChange={setPage}
				/>
			</div>
			<Form onSubmit={handleSubmit}>
				<Table responsive striped bordered hover className="mt-3">
					<thead>
						<tr>
							<th id="no" style={{ fontSize: "1vw" }} width={70}>No</th>
							<th id="ppk" style={{ fontSize: "1vw" }} width={170}>PPK</th>
							<th id="name" style={{ fontSize: "1vw" }} width={190}>Name</th>
							<th id="channeling_company" style={{ fontSize: "1vw" }} width={130}>Channeling Company</th>
							<th id="drawdown_date" style={{ fontSize: "1vw" }} width={190}>Drawdown Date</th>
							<th id="loan_amount" style={{ fontSize: "1vw" }} width={190}>Loan Amount</th>
							<th id="loan_period" style={{ fontSize: "1vw" }} width={190}>Loan Period</th>
							<th id="interest_effective" style={{ fontSize: "1vw" }} width={190}>Interest Eff</th>
							<th id="action" style={{ fontSize: "1vw" }} width={190}>action</th>
						</tr>
					</thead>
					<tbody>
						{recordsJSX}
					</tbody>
				</Table>
				<Button variant="primary" type="submit">
					Submit
				</Button>
			</Form>
		</>
	)
}
