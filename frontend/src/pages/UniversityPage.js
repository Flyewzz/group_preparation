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

    this.currPage = 1;
    this.pageCount = 1;
    this.subjectsService = new SubjectsService();
  }

  subjects = [];

  componentDidMount() {
    const page = 1; // TODO get from url
    this.getSubjects(page);
  }

  getSubjects = (page, name, semester) => {
    const id = this.props.id;
    this.currPage = page;
    this.subjectsService.getPage(id, page, name, semester).then((result) => {
        runInAction(() => {
          this.pageCount = result.pages;
          this.subjects = result.payload ? result.payload : [];
        });
      },
      (error) => {
        console.log(error);
      });
  };

  onPageClick = (event, page) => {
    this.getSubjects(page);
  };

  onSemesterChange = (event) => {
    this.semester = event.target.value;
    this.getSubjects(1, this.name, this.semester);
  };

  render() {
    return (
      <ListContainer title={'Предметы'}
                     subheader={<Filter onSemesterChange={this.onSemesterChange}/>}
                     currPage={this.currPage}
                     pageCount={this.pageCount}
                     onChange={this.onPageClick}
                     items={this.subjects.map((value) =>
                       <Subject subject={value}/>
                     )}/>
    );
  }
}

decorate(UniversityPage, {
  pageCount: observable,
  subjects: observable,
});

export default observer(UniversityPage);
