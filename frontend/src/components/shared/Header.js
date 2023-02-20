import './Header.css';
import { Link } from 'react-router-dom';
import { ICONS_DIR } from '../../const';
import { useStore } from '../../Context';

export default function Header() {
	const { state, dispatch } = useStore();

	return (
		<nav class="navbar navbar-expand-lg bg-light sticky-top">
			<div class="container">
				<Link to="/" className="navbar-brand">
					<img src={`${ICONS_DIR}/sinarmas logo.png`} alt="sinarmas logo" style={{ width: "100px", height: "25px" }} className="d-inline-block align-text-top" />
					<span className="brand-title">Pengajuan Kredit</span>
				</Link>
				<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>
				<div class="collapse navbar-collapse" id="navbarSupportedContent">
					{/* <ul class="navbar-nav me-auto mb-2 mb-lg-0">
						<li class="nav-item">
							<Link to="/" className="nav-link" aria-current="page">Home</Link>
						</li>
					</ul> */}
					<ul class="navbar-nav ms-auto mb-2 mb-lg-0">
						<li class="nav-item dropdown">
							<a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
								{state.user.username}
							</a>
							<ul class="dropdown-menu">
								<li><Link to="#" className="dropdown-item">Action</Link></li>
								<li><hr class="dropdown-divider" /></li>
								<li><Link to="#" className="dropdown-item" onClick={() => dispatch({ type: 'delete' })}>Log out</Link></li>
							</ul>
						</li>
					</ul>
				</div>
			</div>
		</nav>
	)
}
