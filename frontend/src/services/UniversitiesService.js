import config from '../config';

class UniversitiesService {
  getByName = async (name) => {
    const url = config.apiUrl + 'universities?name' + name;
    const options = {method: 'GET', credentials: 'include'};
    const request = new Request(url, options);
    return await fetch(request);
  };

  getPage = async (pageNumber) => {
    const url = config.apiUrl + 'universities?page=' + pageNumber;
    const options = {method: 'GET', credentials: 'include'};
    const request = new Request(url, options);
    return await fetch(request);
  };
}

export default UniversitiesService;