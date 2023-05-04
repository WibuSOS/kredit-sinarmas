import { DRAWDOWN_REPORT_URL, RUPIAH } from "../../../const";
import { useStore } from "../../../Context";
import Swal from 'sweetalert2';
import { StatusCodes, getReasonPhrase } from 'http-status-codes';
import { useEffect, useRef, useState } from 'react';
import Table from 'react-bootstrap/Table';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Pagination, { bootstrap5PaginationPreset } from 'react-responsive-pagination';

export default function DrawdownReport() {
	const { dispatch } = useStore();
	const handleLogOut = () => dispatch({ type: 'logout' });
	const [records, setRecords] = useState([]);
	const [countRecord, setCountRecord] = useState(0);
	const [countPage, setCountPage] = useState(0);
	const [companies, setCompanies] = useState([]);
	const [branches, setBranches] = useState([]);
	const [approvalStatuses, setApprovalStatuses] = useState([]);
	const [page, setPage] = useState(1);
	const limit = useRef(5);
	const filterCompany = useRef('');
	const filterBranch = useRef('');
	const filterStartDate = useRef('');
	const filterEndDate = useRef('');
	const filterApprovalStatus = useRef('');

	useEffect(() => getResource(
		page,
		limit.current,
		filterCompany.current,
		filterBranch.current,
		filterStartDate.current,
		filterEndDate.current,
		filterApprovalStatus.current
	), [page, limit.current]);

	const getResource = (page, limit, company, branch, startDate, endDate, approvalStatus) => {
		console.log("masuk getresource");
		console.log("page:", page);
		fetch(`${DRAWDOWN_REPORT_URL}?page=${page}&limit=${limit}&company=${company}&branch=${branch}&start_date=${startDate}&end_date=${endDate}&approval_status=${approvalStatus}`, {
			method: 'GET',
			credentials: 'include',
			headers: {
				'Accept': 'application/json'
			},
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
				console.log("records:", json.data.records);
				console.log("companies:", json.data.companies);
				console.log("branches:", json.data.branches);
				console.log("count record:", json.data.count_record);
				console.log("count page:", json.data.count_page);
				setRecords(json.data.records ? json.data.records : []);
				setCompanies(json.data.companies ? json.data.companies : []);
				setBranches(json.data.branches ? json.data.branches : []);
				setApprovalStatuses(json.data.approval_statuses ? json.data.approval_statuses : []);
				setCountRecord(json.data.count_record);
				setCountPage(json.data.count_page);
			})
			.catch(err => console.error(err));
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		console.log("submitted");
		console.log("approval:", filterApprovalStatus.current);
		console.log("page:", page);
		if (page === 1) {
			getResource(
				page,
				limit.current,
				filterCompany.current,
				filterBranch.current,
				filterStartDate.current,
				filterEndDate.current,
				filterApprovalStatus.current
			);
		} else {
			setPage(1);
		}
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
			</tr>
		)
	});

	const companiesJSX = companies.map((company, index) => (
		<option key={index + 1} value={company.company_short_name}>{`${company.company_code} - ${company.company_short_name}`}</option>
	));
	const branchesJSX = branches.map((branch, index) => (
		<option key={index + 1} value={branch.code}>{`${branch.code} - ${branch.description}`}</option>
	));
	const approvalStatusesJSX = approvalStatuses.map((approvalStatus, index) => (
		<option key={index + 1} value={approvalStatus}>{approvalStatus}</option>
	));

	return (
		<>
			<Form onSubmit={handleSubmit} className="mt-4">
				<Row>
					<Col xl="auto">
						<Form.Label>Company</Form.Label>
						<Form.Select aria-label="filter company" defaultValue={filterCompany.current} onChange={(e) => filterCompany.current = e.target.value}>
							<option key={0} value="">Pilih company</option>
							{companiesJSX}
						</Form.Select>
					</Col>
					<Col xl="auto">
						<Form.Label>Branch</Form.Label>
						<Form.Select aria-label="filter branch" defaultValue={filterBranch.current} onChange={(e) => filterBranch.current = e.target.value}>
							<option key={0} value="">Pilih branch</option>
							{branchesJSX}
						</Form.Select>
					</Col>
					<Col xl="auto">
						<Form.Label>Start Date</Form.Label>
						<Form.Control type="date" aria-label="filter start date" onChange={(e) => filterStartDate.current = e.target.value} />
					</Col>
					<Col xl="auto">
						<Form.Label>End Date</Form.Label>
						<Form.Control type="date" aria-label="filter end date" onChange={(e) => filterEndDate.current = e.target.value} />
					</Col>
					<Col xl="auto">
						<Form.Label>Approval Status</Form.Label>
						<Form.Select aria-label="filter approval status" defaultValue={filterApprovalStatus.current} onChange={(e) => filterApprovalStatus.current = e.target.value}>
							<option key={0} value="">Pilih approval status</option>
							{approvalStatusesJSX}
						</Form.Select>
					</Col>
					<Col xl="auto">
						<Button variant="primary" type="submit" className="mt-4">
							Submit
						</Button>
					</Col>
				</Row>
			</Form>
			<div className="mt-5">
				<Pagination
					{...bootstrap5PaginationPreset}
					current={page}
					total={countPage}
					onPageChange={setPage}
				/>
			</div>
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
					</tr>
				</thead>
				<tbody>
					{recordsJSX}
				</tbody>
			</Table>
		</>
	)
}
