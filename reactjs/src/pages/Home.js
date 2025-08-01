import { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import Alert from 'react-bootstrap/Alert';
import ModalComponent from '../components/ModalComponent'
import { upload } from '../services/Request'
import { setMessageState } from '../utils/UIUtils'
import { Constants } from '../utils/Constants';

const Home = () => {

    const [selectedFile, setSelectedFile] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const [message, setMessage] = useState({
        label: '',
        variant: ''
    });

    const uploadFile = () => {
        setShowModal(true);
        upload(selectedFile)
            .then((response) => {
                setMessage(() => setMessageState(response, Constants.ALERT_SUCCESS));
            }).catch((error) => {
                setMessage(() => setMessageState(error, Constants.ALERT_DANGER));
            }).finally(() => {
                setShowModal(false);
            });
    };

    return (
        <>
            <Alert variant={message.variant} hidden={message.label === null}>{message.label}</Alert>
            <Form.Group controlId="formFile" className="mb-3">
                <Form.Control type="file" accept="image/png, image/jpeg" onChange={(event) => {
                    setMessage(() => setMessageState('', ''));
                    const file = event.target.files[0];
                    const maxSize = 2 * 1024 * 1024;
                    if (file.size > maxSize) {
                        setMessage(() => setMessageState(Constants.MSG_FILE_EXCEEDS, Constants.ALERT_WARNING));
                    } else {
                        setSelectedFile(file);
                    }
                }} />
                <br />
                <Button className="button" variant="primary" disabled={!selectedFile} onClick={uploadFile}>{Constants.UPLOAD}</Button>
            </Form.Group>
            <ModalComponent showModal={showModal} label={Constants.MODAL_LABEL_UPLOAD}/>
        </>
    )
}

export default Home