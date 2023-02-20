import { useState } from "react";
import { Button, Offcanvas } from "react-bootstrap";
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
					<Offcanvas.Title>Offcanvas</Offcanvas.Title>
				</Offcanvas.Header>
				<Offcanvas.Body>
					Some text as placeholder. In real life you can have the elements you
					have chosen. Like, text, images, lists, etc.
				</Offcanvas.Body>
			</Offcanvas>
		</div>
	)
}
