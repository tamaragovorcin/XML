import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class ModalDialog extends Component {
  render() {
    return (
      <Modal
        show={this.props.show}
        size="lg"
        aria-labelledby="contained-modal-title-vcenter"
        centered
        onHide={this.props.onCloseModal}
      >
        <Modal.Header closeButton>
          <Modal.Title id="contained-modal-title-vcenter">
            {this.props.header}
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <h4>{this.props.subheader}</h4>
          <p>{this.props.text}</p>
        </Modal.Body>
        <Modal.Footer>
          <Button onClick={this.props.onCloseModal} href={this.props.href}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    );
  }
}

export default ModalDialog;
