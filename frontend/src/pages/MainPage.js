import React from "react";
import University from "../components/universities/University";
import ListContainer from "../components/common/ListContainer"
import {observer} from "mobx-react";
import UniversitiesService from "../services/UniversitiesService";
import {decorate, observable, runInAction} from "mobx";

class MainPage extends React.Component {
  constructor(props) {
    super(props);

    this.currPage = 1;
    this.pageCount = 1;
    this.universitiesService = new UniversitiesService();
  }

  universities = [];

  componentDidMount() {
    const page = 1; // TODO get from url
    this.getUniversities(page);
  }

  getUniversities = (page) => {
    this.currPage = page;
    this.universitiesService.getPage(page).then((result) => {
        console.log(result);
        runInAction(() => {
          this.pageCount = result.pages;
          this.universities = result.payload;
        });
      },
      (error) => {
        console.log(error);
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
