import React from "react";
import University from "../components/universities/University";
import ListContainer from "../components/common/ListContainer"
import {observer} from "mobx-react";
import UniversitiesService from "../services/UniversitiesService";
import {decorate, observable, runInAction} from "mobx";

const data = [
  {
    id: 0,
    short: 'МГТУ им. Баумана',
    full: 'Московский государственный технический университет имени Баумана'
  },
  {
    id: 1,
    short: 'ВШЭ',
    full: 'Высшая школа экономика'
  },
  {
    id: 2,
    short: 'МГУ им. Ломоносова',
    full: 'Московский государственный университет имени Ломоносова'
  },
  {
    id: 3,
    short: 'ГУУ',
    full: 'Государственный университет управления'
  },
  {
    id: 4,
    short: 'МГТУ им. Баумана',
    full: 'Московский государственный технический университет имени Баумана'
  },
  {
    id: 5,
    short: 'ВШЭ',
    full: 'Высшая школа экономика'
  },
  {
    id: 6,
    short: 'МГУ им. Ломоносова',
    full: 'Московский государственный университет имени Ломоносова'
  },
  {
    id: 7,
    short: 'ГУУ',
    full: 'Государственный университет управления'
  },
  {
    id: 8,
    short: 'МГУ им. Ломоносова',
    full: 'Московский государственный университет имени Ломоносова'
  },
];

class MainPage extends React.Component {
  constructor(props) {
    super(props);

    this.pageCount = 1;
    this.universitiesService = new UniversitiesService();
  }

  universities = [];

  componentDidMount() {
    const page = 1; // get from url
    this.getUniversities(page);
  }

  getUniversities = (page) => {
    this.currPage = page;
    this.universitiesService.getPage(page).then((result) => {
        runInAction(() => {
          // this.pageCount = result.pages;
          // this.universities = result.payload;
          this.pageCount = 5;
          this.universities = data;
        });
      },
      (error) => {
        console.log(error);
        this.pageCount = 5;
        this.universities = data;
      });
  };

  onPageClick = (event, page) => {
    this.getUniversities(page);
  };

  render() {
    return (
      <ListContainer title={'Университеты'}
                     currPage={this.currPage}
                     pageCount={this.pageCount}
                     onChange={this.onPageClick}
                     items={this.universities.map((value) =>
                       <University key={value.id} university={value}/>
                     )}/>
    );
  }
}

decorate(MainPage, {
  pageCount: observable,
  universities: observable,
});

export default observer(MainPage);
