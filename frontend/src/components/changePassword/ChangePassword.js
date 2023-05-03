import { CHANGE_PASSWORD_URL } from "../../const";
import Swal from 'sweetalert2';
import { ReasonPhrases, StatusCodes, getReasonPhrase, getStatusCode } from 'http-status-codes';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { useStore } from "../../Context";

export default function ChangePassword() {
	const { state, dispatch } = useStore();
	const handleLogOut = () => dispatch({ type: 'logout' });

	const handleSubmit = (e) => {
		e.preventDefault();
		const formData = new FormData(e.currentTarget);
		const body = {
			old_password: formData.get("old_password"),
			new_password: formData.get("new_password"),
			confirm_password: formData.get("confirm_password")
		};
		if (!body.old_password || !body.new_password || !body.confirm_password) {
			Swal.fire({ icon: 'error', title: 'Error', text: 'Form Harus Terisi' });
			return
		}
		if (body.new_password.length != 8) {
			Swal.fire({ icon: 'error', title: 'Error', text: 'Password Baru Harus 8 Karakter' });
			return
		}
		if (body.new_password != body.confirm_password) {
			Swal.fire({ icon: 'error', title: 'Error', text: 'Pastikan Password Baru Benar' });
			return
		}
		fetch(CHANGE_PASSWORD_URL, {
			method: 'PATCH',
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json',
				'Authorization': state.user.token
			},
			body: JSON.stringify(body)
		})
			.then(res => {
				if (!res.ok && (res.status == StatusCodes.UNAUTHORIZED || res.status == StatusCodes.FORBIDDEN)) {
					handleLogOut();
					throw new Error(`${res.status}::${res.statusText}`);
				}
				return res.json()
			})
			.then(json => {
				if (json.code != StatusCodes.OK) {
					Swal.fire({ icon: 'error', title: 'Error', text: getReasonPhrase(json.code) });
					throw new Error(`${json.code}::${getReasonPhrase(json.code)}::${json.message}`);
				}
				Swal.fire({ icon: 'success', title: 'Success', showConfirmButton: false, timer: 1500 });
				handleLogOut();
			})
			.catch(err => console.error(err));
	};

	return (
		<Form onSubmit={handleSubmit} className="mt-3">
			<Form.Group className="mb-3">
				<Form.Label>Old Password</Form.Label>
				<Form.Control type="password" name="old_password" placeholder="Old password" />
			</Form.Group>
			<Form.Group className="mb-3">
				<Form.Label>New Password</Form.Label>
				<Form.Control type="password" name="new_password" placeholder="New password" />
			</Form.Group>
			<Form.Group className="mb-3">
				<Form.Label>Confirm Password</Form.Label>
				<Form.Control type="password" name="confirm_password" placeholder="Confirm password" />
			</Form.Group>
			<Button variant="primary" type="submit">
				Submit
			</Button>
		</Form>
	)
}
