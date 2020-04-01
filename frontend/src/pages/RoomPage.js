import React from "react";
import {observer} from "mobx-react";
import UniversitiesService from "../services/UniversitiesService";
import {decorate, observable, runInAction} from "mobx";
import MessageList from "../components/chat/MessageList";
import RoomPathHeader from "./RoomPathHeader";
import SubjectsService from "../services/SubjectsService";
import RoomService from "../services/RoomsService";
import RoomTabs from "../components/rooms/RoomTabs";

class RoomPage extends React.Component {
  constructor(props) {
    super(props);

    this.roomService = new RoomService();
    this.subjectService = new SubjectsService();
    this.universityService = new UniversitiesService();
  }

  university = {name: ''};
  subject = {name: ''};
  room = {name: ''};

  componentDidMount() {
  }

  getRoom = () => {
    const id = this.props.id;
    this.roomService.getById(id).then((result) => {
        runInAction(() => {
          this.room = result;
        });
        this.getSubject();
      },
      (error) => {
        console.log(error)
      })
  };

  getSubject = () => {
    const id = this.room.subject_id;
    this.subjectService.getById(id).then((result) => {
        runInAction(() => {
          this.subject = result;
        });
        this.getUniversity();
      },
      (error) => {
        console.log(error)
      })
  };

  getUniversity = () => {
    const id = this.subject.university_id;
    this.universityService.getById(id).then((result) => {
        runInAction(() => {
          this.university = result;
        });
      },
      (error) => {
        console.log(error)
      })
  };

  render() {
    return (
      <>
        <RoomPathHeader university={this.university}
                        subject={this.subject}
                        roomName={this.room.name}/>
        <RoomTabs/>
        <MessageList/>
      </>
    );
  }
}
// fuckfu
decorate(RoomPage, {
  pageCount: observable,
  university: observable,
  subject: observable,
});

export default observer(RoomPage);
