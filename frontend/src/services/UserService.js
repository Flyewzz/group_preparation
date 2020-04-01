import config from '../config';

class UserService {
  signIn = async () => {
    const email = 'emailemail';
    const password = '123';
    const url = config.apiUrl + `signin?email=${email}&password=${password}`;
    const options = {method: 'POST', credentials: 'include'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default UserService;