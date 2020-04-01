import React from "react";
import {observable, runInAction, decorate} from "mobx"
import {observer} from "mobx-react"
import Subject from "../components/subjects/Subject";
import Filter from "../components/subjects/Filter";
import SubjectsService from "../services/SubjectsService";
import SubjectPathHeader from "../components/subjects/SubjectPathHeader";
import UniversitiesService from "../services/UniversitiesService";
import SemesterFilter from "../components/subjects/SemesterFilter";
import Container from "@material-ui/core/Container"
import Typography from "@material-ui/core/Typography";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Divider from "@material-ui/core/Divider"
import Pagination from "@material-ui/lab/Pagination"
import {makeStyles} from "@material-ui/core/styles";


const useStyles = makeStyles(() => ({
  wrapper: {
    display: 'flex',
  },
}));

function Wrapper(props) {
  const classes = useStyles();
  return (
    <div className={classes.wrapper}>
      {props.children}
    </div>
  )
}

const listUseStyles = makeStyles(() => ({
  wrapper: {
    width: '95%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '0 auto',
    padding: '0 10pt',
  },
  header: {
    fontWeight: 'bold',
    marginBottom: '5pt',
  },
  list: {
    width: '95%',
  },
}));

function formList(items) {
  const result = [];
  items.forEach((value, key) => {
    result.push(<ListItem button key={key}>
      {value}
    </ListItem>);

    if (key !== items.length - 1) {
      result.push(
        <Divider key={items.length + key}
                 variant="middle" component="li"/>
      );
    }
  });

  return result;
}

function ListContainer(props) {
  const styles = listUseStyles();

  return (
    <Container maxWidth="lg" className={styles.wrapper}>
      <List className={styles.list} subheader={props.subheader}>
        {formList(props.items)}
      </List>
      <Pagination count={props.pageCount}
                  page={props.currPage}
                  onChange={props.onChange}
                  size="large"
                  shape="rounded"
                  color="primary"/>
    </Container>
  );
}

class UniversityPage extends React.Component {
  constructor(props) {
    super(props);

    this.currPage = 1;
    this.pageCount = 1;
    this.subjectsService = new SubjectsService();
    this.universityService = new UniversitiesService();
  }

  university = {};
  subjects = [];

  componentDidMount() {
    const page = 1; // TODO get from url
    this.getUniversity();
    this.getSubjects(page);
  }

  getUniversity = () => {
    const id = this.props.id;
    this.universityService.getById(id).then((result) => {
      runInAction(() => {
        this.university = result;
      });
    });
  };

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
    if (this.semester === 'any') {
      this.semester = null;
    }
    this.getSubjects(1, this.name, this.semester);
  };

  onNameChange = (event) => {
    this.name = event.target.value;
    this.getSubjects(1, this.name, this.semester);
  };

  render() {
    return (
      <>
        <SubjectPathHeader university={this.university}/>
        <Wrapper>
          <SemesterFilter onSemesterChange={this.onSemesterChange}/>
          <ListContainer subheader={<Filter onNameChange={this.onNameChange}/>}
                         currPage={this.currPage}
                         pageCount={this.pageCount}
                         onChange={this.onPageClick}
                         items={this.subjects.map((value) =>
                           <Subject subject={value}/>
                         )}/>
        </Wrapper>
      </>
    );
  }
}

decorate(UniversityPage, {
  pageCount: observable,
  university: observable,
  subjects: observable,
});

export default observer(UniversityPage);
