import { Button, Col, Row, Container, Card, Form, InputGroup } from 'react-bootstrap';
import Swal from 'sweetalert2';
import './Login.css';
import { useStore } from '../../Context';
import { ICONS_DIR, LOGIN_URL } from '../../const';

export default function Login() {
	const { dispatch } = useStore();

	const handleSubmit = async e => {
		e.preventDefault();
		const formData = new FormData(e.currentTarget);
		const body = {
			username: formData.get("username"),
			password: formData.get("password")
		};
		console.log(body.username);
		console.log(body.password);
		fetch(LOGIN_URL, {
			method: 'POST',
			credentials: 'include',
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(body)
		})
			.then(res => res.json())
			.then(json => {
				if (!json.data) {
					Swal.fire({ icon: 'error', title: 'Login Gagal', text: 'username/password salah' });
					return
				}
				Swal.fire({ icon: 'success', title: 'Login Berhasil', showConfirmButton: false, timer: 1500 });
				const token = json.data?.token;
				dispatch({ type: 'set', payload: { token } });
			})
			.catch(err => console.error(err));
	};

	return (
		<Container fluid as='main' className='login-body'>
			<Col className='d-flex justify-content-center align-items-center mt-8-percent mb-3'>
				<div className='my-5 mx-auto log-shadow text-center p-5'>
					<Card.Body className='w-100 d-flex flex-column'>
						<Form onSubmit={e => handleSubmit(e)}>
							<Row>
								<InputGroup className="btn-shadow mb-2 mt-5">
									<InputGroup.Text id='usernameIcon' className='btn-input'><img src={`${ICONS_DIR}/username.png`} alt='usernameIcon' style={{ width: '20px', height: '20px' }} /></InputGroup.Text>
									<Form.Control
										name="username"
										className="btn-input"
										placeholder="Enter username"
										aria-label="username"
										aria-describedby="usernameIcon"
									/>
								</InputGroup>
								<InputGroup className="btn-shadow mb-3">
									<InputGroup.Text id='passwordIcon' className='btn-input'><img src={`${ICONS_DIR}/password.png`} alt='passwordIcon' style={{ width: '20px', height: '20px' }} /></InputGroup.Text>
									<Form.Control
										name="password"
										className="btn-input"
										placeholder="Enter password"
										aria-label="password"
										aria-describedby="passwordIcon"
										type="password"
									/>
								</InputGroup>
							</Row>
							<Row>
								<Button className='btn-shadow btn-input' size="lg" variant='Primary' style={{ backgroundColor: "#128297", color: "white" }} type='submit'>
									Login
								</Button>
							</Row>
							{/* <Row>
								<Link to='/register' className='a-regist mt-2'>
									Register
								</Link>
							</Row> */}
						</Form>
					</Card.Body>
				</div>
			</Col>
		</Container >
	)
}
