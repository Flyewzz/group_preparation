import React from "react";
import RoomsList from "./RoomsList";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import RoomsService from "../../services/RoomsService";

const testData = [
  {
    id: 1,
    name: 'РК №1 Решение дифференциальных уравнений',
    uuid: 'aaa',
    author: 'username',
    type: 'ЛР',
    date: '01.01.2020',
    banned: false,
  },
  {
    id: 1,
    name: 'РК №1 Решение дифференциальных уравнений',
    uuid: 'aaa',
    author: 'username',
    type: 'ЛР',
    date: '01.01.2020',
    banned: false,
  },
  {
    id: 1,
    name: 'РК №1 Решение дифференциальных уравнений',
    uuid: 'aaa',
    author: 'username',
    type: 'ЛР',
    date: '01.01.2020',
    banned: false,
  },
];

class RoomsTab extends React.Component {
  constructor(props) {
    super(props);

    this.currPage = 1;
    this.pageCount = 1;
    this.roomsService = new RoomsService();
  }

  rooms = [];

  componentDidMount() {
    const page = 1; // TODO get from url
    this.getRooms(page);
  }

  getRooms = (page, name, type) => {
    this.currPage = page;

    this.roomsService.getPage(page, name, type).then((result) => {
        runInAction(() => {
          this.pageCount = result.pages;
          this.pageCount = 1; // TODO temporary
          this.rooms = result.payload ? result.payload : [];
        });
      },
      (error) => {
        this.rooms = testData;
        console.log(error);
      });
  };

  onPageClick = (event, page) => {
    this.getRooms(page);
  };

  render() {
    return (
        <RoomsList id={this.props.id}
                   data={this.rooms}
                   pageCount={this.pageCount}
                   currPage={this.currPage}
                   onChange={this.onPageClick}
                   onNameChange={this.props.onNameChange}
                   onTypeChange={this.props.onTypeChange}/>
    );
  }
}

decorate(RoomsTab, {
  pageCount: observable,
  rooms: observable,
});

export default observer(RoomsTab);