import { useState } from "react";
import { Button, Offcanvas } from "react-bootstrap";
import { Link } from "react-router-dom";
import { ICONS_DIR } from "../../const";
import "./Sidebar.css";

export default function Sidebar() {
	const [show, setShow] = useState(false);
	const handleShow = () => setShow(true);
	const handleClose = () => setShow(false);

	return (
		<div id="sidebar-app">
			<div onClick={handleShow} id="sidebar-app-toggler">
				<div id="sidebar-app-toggler-icon"><img src={`${ICONS_DIR}/chevron-right-30.png`} alt="chevronRight" style={{ width: "12px", height: "12px" }} /></div>
				<div>Sidebar</div>
			</div>

			<Offcanvas show={show} onHide={handleClose}>
				<Offcanvas.Header closeButton>
					<Offcanvas.Title>Main Menu</Offcanvas.Title>
				</Offcanvas.Header>
				<Offcanvas.Body>
					<Link to="/" className="sidebar-app-menu-body-link" onClick={handleClose}>
						<div className="sidebar-app-menu-body-item">
							Home
						</div>
					</Link>
					<hr />
					<Link to="/" className="sidebar-app-menu-body-link" onClick={handleClose}>
						<div className="sidebar-app-menu-body-item">
							Another Menu
						</div>
					</Link>
					<hr />
				</Offcanvas.Body>
			</Offcanvas>
		</div>
	)
}
