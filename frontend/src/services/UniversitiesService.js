import config from '../config';

class UniversitiesService {
  getByName = async (name) => {
    const url = config.apiUrl + 'universities?name' + name;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getById = async (id) => {
    const url = config.apiUrl + `university?id=${id}`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getAll = async () => {
    const url = config.apiUrl + 'universities';
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getPage = async (pageNumber) => {
    const url = config.apiUrl + 'universities?page=' + pageNumber;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default UniversitiesService;