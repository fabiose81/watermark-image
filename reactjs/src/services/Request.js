export function upload(file) {
    return new Promise((resolve, reject) => {
        const apiUrl = process.env.REACT_APP_API_URL;

        const formData = new FormData();
        formData.append('file', file);

        const requestOptions = {
            method: 'POST',
            body: formData
        };

        fetch(apiUrl.concat('/upload'), requestOptions)
         .then((response) => {
                if (response.status === 200) {
                    return response.text()
                } else if (response.status === 500) {
                    //to-do
                    reject('Internal Error')
                } else {
                    reject(response.text())
                }
            })
            .then((result) => {
                resolve(result)
            })
            .catch((error) => 
                reject(error.message));
    })
}