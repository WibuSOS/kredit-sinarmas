import { CHECKLIST_PENCAIRAN_URL } from "../../../const";
import { useStore } from "../../../Context";
import Swal from 'sweetalert2';
import {
	ReasonPhrases,
	StatusCodes,
	getReasonPhrase,
	getStatusCode,
} from 'http-status-codes';
import { useEffect, useRef, useState } from 'react';
import Table from 'react-bootstrap/Table';

export default function ChecklistPencairan() {
	const { state, dispatch } = useStore();
	const handleLogOut = () => dispatch({ type: 'delete' });
	const [records, setRecords] = useState([]);
	const [countRecord, setCountRecord] = useState(0);
	const [countPage, setCountPage] = useState(0);
	const page = useRef(1);
	const limit = useRef(10);

	useEffect(() => getResource(page.current, limit.current), [page.current, limit.current]);

	const getResource = (page, limit) => {
		fetch(`${CHECKLIST_PENCAIRAN_URL}?page=${page}&limit=${limit}`, {
			method: 'GET',
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json',
				'Authorization': state.user.token
			},
			// body: JSON.stringify(body)
		})
			.then(res => {
				if (!res.ok && res.status == StatusCodes.UNAUTHORIZED) {
					handleLogOut();
					throw new Error(`${res.status}::${res.statusText}`);
				}
				if (!res.ok) {
					throw new Error(`${res.status}::${res.statusText}`);
				}
				return res.json()
			})
			.then(json => {
				// if (json.code == StatusCodes.UNAUTHORIZED) {
				// 	handleLogOut();
				// 	return
				// }
				if (!json.data) {
					Swal.fire({ icon: 'error', title: 'Error', text: json.message });
					return
				}
				console.log(json.data);
				setRecords(json.data.records ? json.data.records : []);
				setCountRecord(json.data.count_record);
				setCountPage(json.data.count_page);
			})
			.catch(err => console.error(err));
	};

	const recordsJSX = records.map((record, index) => {
		let no = (page.current - 1) * limit.current + index + 1;
		return (
			<tr key={index}>
				<td headers="no" style={{ fontSize: "1vw" }}>{no}</td>
				<td headers="ppk" style={{ fontSize: "1vw" }}>{record.ppk}</td>
				<td headers="name" style={{ fontSize: "1vw" }}>{record.name}</td>
				<td headers="channeling_company" style={{ fontSize: "1vw" }}>{record.channeling_company}</td>
				<td headers="drawdown_date" style={{ fontSize: "1vw" }}>{record.drawdown_date.substring(0, 10)}</td>
				<td headers="loan_amount" style={{ fontSize: "1vw" }}>{record.loan_amount}</td>
				<td headers="loan_period" style={{ fontSize: "1vw" }}>{record.loan_period}</td>
				<td headers="interest_effective" style={{ fontSize: "1vw" }}>{record.interest_effective}</td>
				{/* <td headers="action" style={{ fontSize: "1vw" }} >
					<div className="form-check">
						<input className="form-check-input" type="checkbox" value="" id="flexCheckDefault" onChange={() => checked(i)} checked={data.check} />
						<label className="form-check-label" htmlFor="flexCheckDefault">
							Pilih
						</label>
					</div>
				</td> */}
			</tr>
		)
	});

	return (
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
	)
}
