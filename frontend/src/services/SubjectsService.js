import config from '../config';

class SubjectsService {
  getPage = async (id, pageNumber, name, semester) => {
    let url = config.apiUrl + 'university/' + id + '/subjects?page=' + pageNumber;
    if (name) {
      url += '&name=' + name;
    }
    if (semester) {
      url += '&semester=' + semester;
    }
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default SubjectsService;