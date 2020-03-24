import config from '../config';

class MaterialService {
  getPage = async (id, pageNumber, name, typeId) => {
    let url = config.apiUrl + 'subject/' + id + '/materials?page=' + pageNumber;
    if (name) {
      url += '&name=' + name;
    }
    if (typeId) {
      url += '&type_id=' + typeId;
    }
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default MaterialService;