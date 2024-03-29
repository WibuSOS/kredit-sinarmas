import './Header.css';
import { Link } from 'react-router-dom';
import { ICONS_DIR, LOGIN_URL } from '../../const';
import { useStore } from '../../Context';
import Swal from 'sweetalert2';

export default function Header() {
	const { state, dispatch } = useStore();
	const handleLogOut = () => {
		fetch(LOGIN_URL, {
			method: 'DELETE',
			credentials: 'include',
			headers: {
				'Accept': 'application/json'
			}
		})
			.then(res => {
				if (!res.ok) {
					Swal.fire({ icon: 'error', title: 'Something is Wrong', text: 'Proceeding to logout\nRedirected to Login', showConfirmButton: false, timer: 1500 });
					dispatch({ type: 'logout' });
					throw new Error(`${res.status}::${res.statusText}`);
				}
				Swal.fire({ icon: 'success', title: 'Logout Success', text: 'Redirected to Login', showConfirmButton: false, timer: 1500 });
				dispatch({ type: 'logout' });
			})
			.catch(err => console.error(err));
	};

	return (
		<nav className="navbar navbar-expand-lg bg-light sticky-top">
			<div className="container">
				<Link to="/" className="navbar-brand">
					<img src={`${ICONS_DIR}/sinarmas logo.png`} alt="sinarmas logo" style={{ width: "100px", height: "25px" }} className="d-inline-block align-text-top" />
					<span className="brand-title">Pengajuan Kredit</span>
				</Link>
				<button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
					<span className="navbar-toggler-icon"></span>
				</button>
				<div className="collapse navbar-collapse" id="navbarSupportedContent">
					{/* <ul className="navbar-nav me-auto mb-2 mb-lg-0">
						<li className="nav-item">
							<Link to="/" className="nav-link" aria-current="page">Home</Link>
						</li>
					</ul> */}
					<ul className="navbar-nav ms-auto mb-2 mb-lg-0">
						<li className="nav-item dropdown">
							<a className="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
								{state.user.name}
							</a>
							<ul className="dropdown-menu">
								<li><Link to="/change_password" className="dropdown-item">Change password</Link></li>
								<li><hr className="dropdown-divider" /></li>
								<li><Link to="#" className="dropdown-item" onClick={handleLogOut}>Log out</Link></li>
							</ul>
						</li>
					</ul>
				</div>
			</div>
		</nav>
	)
}
