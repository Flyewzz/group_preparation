import React from "react";
import {observable, runInAction, decorate} from "mobx"
import {observer} from "mobx-react"
import ListContainer from "../components/common/ListContainer"
import Subject from "../components/subjects/Subject";
import Filter from "../components/subjects/Filter";
import SubjectsService from "../services/SubjectsService";

const data = [
  {
    id: 1,
    name: 'Физика',
    files: 2,
  },
  {
    id: 2,
    name: 'Математический анализ',
    files: 2,
  },
  {
    id: 3,
    name: 'Дискретная математика',
    files: 2,
  },
  {
    id: 4,
    name: 'Экология',
    files: 2,
  },
  {
    id: 5,
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    id: 6,
    name: 'Физика',
    files: 2,
  },
  {
    id: 7,
    name: 'Математический анализ',
    files: 2,
  },
  {
    id: 8,
    name: 'Дискретная математика',
    files: 2,
  },
  {
    id: 9,
    name: 'Экология',
    files: 2,
  },
  {
    id: 10,
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    id: 11,
    name: 'Аналитическая геометрия',
    files: 2,
  },
  {
    id: 12,
    name: 'Аналитическая геометрия',
    files: 2,
  },
];

class UniversityPage extends React.Component {
  constructor(props) {
    super(props);

    this.pageCount = 1;
    this.subjectsService = new SubjectsService();
  }

  subjects = [];

  componentDidMount() {
    const page = 1; // get from url
    console.log('mounting');
    this.getSubjects(page);
  }

  getSubjects = (page) => {
    this.currPage = page;
    this.subjectsService.getPage(page).then((result) => {
        runInAction(() => {
          // this.pageCount = result.pages;
          // this.universities = result.payload;
          this.pageCount = 5;
          this.subjects = data;
        });
      },
      (error) => {
        console.log(error);
        this.pageCount = 5;
        this.subjects = data;
      });
  };

  onPageClick = (event, page) => {
    this.getSubjects(page);
  };

  render() {
    return (
      <ListContainer title={'Предметы'}
                     subheader={<Filter/>}
                     currPage={this.currPage}
                     pageCount={this.pageCount}
                     onChange={this.onPageClick}
                     items={this.subjects.map((value) =>
                       <Subject subject={value}/>
                     )}>
      </ListContainer>
    );
  }
}

decorate(UniversityPage, {
  pageCount: observable,
  subjects: observable,
});

export default observer(UniversityPage);
