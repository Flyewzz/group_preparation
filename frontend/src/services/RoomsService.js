import config from '../config';

class RoomsService {
  getPage = async (pageNumber, name, typeId) => {
    let url = config.apiUrl + 'rooms';
    if (name) {
      url += '&name=' + name;
    }
    if (typeId) {
      url += '&type_id=' + typeId;
    }
    const options = {method: 'GET', credentials: 'include'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getById = async (id) => {
    const url = config.apiUrl + 'room?id=' + id ;
    const options = {method: 'GET', credentials: 'include'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default RoomsService;