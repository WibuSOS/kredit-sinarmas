import React, { useState } from 'react';
import { Button, Image, Col, Row, Container, Card, Form, InputGroup } from 'react-bootstrap';
import Swal from 'sweetalert2';
import './Login.css';
import { useStore } from '../../Context';
import { LOGIN_URL } from '../../const';

export default function Login() {
	const { dispatch } = useStore();
	// const [username, setUsername] = useState();
	// const [password, setPassword] = useState();

	// const handleUsername = e => {
	// 	const { value } = e.target;
	// 	setUsername(value);
	// };

	// const handlePassword = e => {
	// 	const { value } = e.target;
	// 	setPassword(value);
	// };

	const handleSubmit = async e => {
		e.preventDefault();
		const formData = new FormData(e.currentTarget);
		const body = {
			username: formData.get("username"),
			password: formData.get("password")
		};
		fetch(LOGIN_URL, {
			method: 'POST',
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
				const id = json.data?.id;
				const username = json.data?.username;
				const name = json.data?.name;
				const token = json.data?.token;
				localStorage.setItem('id', id);
				localStorage.setItem('username', username);
				localStorage.setItem('name', name);
				localStorage.setItem('token', token);
				dispatch({ type: 'set', payload: { id, username, name, token } });
			})
			.catch(err => console.error(err));
	};

	return (
		<div className='login-body'>
			<Container fluid className='login-container mb-3'>
				<Row className='d-flex justify-content-center align-items-center'>
					<Col col='12'>
						<div className='my-5 mx-auto log-shadow text-center p-5'>
							<Image src="assets/pega-logo.svg" width="250" height="50" />
							<Card.Body className='w-100 d-flex flex-column'>
								<Row>
									<InputGroup className="btn-shadow mb-2 mt-5" onChange={this.handleUsername} value={this.state.username}>
										<InputGroup.Text id="basic-addon1" className='btn-input'><img src={ICONS + "user2.png"} alt={"dd"} style={{ width: "20px", height: "20px", }} /></InputGroup.Text>
										<Form.Control
											className='btn-input'
											placeholder="Username"
											aria-label="Username"
											aria-describedby="basic-addon1"
										/>
									</InputGroup>
									<InputGroup className="btn-shadow mb-3" onChange={this.handlePassword} value={this.state.username}>
										<InputGroup.Text className='btn-input' id="basic-addon1"><img src={ICONS + "lock.png"} alt={"dd"} style={{ width: "20px", height: "20px", }} /></InputGroup.Text>
										<Form.Control
											className='btn-input'
											placeholder="Username"
											aria-label="Username"
											aria-describedby="basic-addon1"
											type="password"
										/>
									</InputGroup>
								</Row>
								<Row className='mt-1'>
									<Button className='btn-shadow btn-input' size="lg" variant='Primary' style={{ backgroundColor: "#128297", color: "white" }} onClick={(e) => this.handleSubmit(e)}>
										Login
									</Button>
								</Row>
								<Row>
									<a size="lg" href='/register' className='a-regist mt-2'>
										Register
									</a>
								</Row>
							</Card.Body>
						</div>
					</Col>
				</Row>
			</Container>
		</div>
	)
}
