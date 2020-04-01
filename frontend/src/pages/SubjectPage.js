import React from "react";
import Container from "@material-ui/core/Container";
import {makeStyles} from "@material-ui/core/styles";
import FilterLine from "../components/material/FilterLine";
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import FolderRoundedIcon from '@material-ui/icons/FolderRounded';
import SearchIcon from '@material-ui/icons/Search';
import GroupRoundedIcon from '@material-ui/icons/GroupRounded';
import AddIcon from '@material-ui/icons/Add'
import FormatListNumberedRoundedIcon from '@material-ui/icons/FormatListNumberedRounded';
import SubjectsService from "../services/SubjectsService";
import MaterialService from "../services/MaterialService";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import MaterialTab from "../components/material/MaterialTab";
import MaterialsPathHeader from "../components/material/MaterialsPathHeader";
import UniversitiesService from "../services/UniversitiesService";
import TypeFilter from "../components/material/TypeFilter";
import InputBase from "@material-ui/core/InputBase";
import Button from "@material-ui/core/Button";
import RoomsTab from "../components/rooms/RoomsTab";
import AddDialog from "../components/rooms/AddDialog";

const useStyles = makeStyles((theme) => ({
  wrapper: {
    margin: '0',
  },
  tabPanelWrapper: {
    width: '100%',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'start',
  },
  highPanel: {
    width: '84%'
  },
  tabPanel: {
    width: '100%',
  },
  left: {
    marginLeft: '55pt',
    width: 'max-content'
  },
  hr: {
    color: '#efffff',
    transform: 'translateY(-1px)',
    margin: 0,
  },
  indicator: {
    zIndex: 10000,
  },
  root: {
    backgroundColor: theme.palette.background.paper,
    marginLeft: '55pt',
    display: 'flex',
  },
  flex: {
    display: 'flex',
    flexDirection: 'row',
    width: '100%',
  },
  tabs: {
    borderRight: `1px solid ${theme.palette.divider}`,
    border: '1px solid black',
    borderRadius: '6px 0 0 6px',
    minWidth: '0',
  },
  tab: {
    minWidth: 0,
  },
  label: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'center',
  },
  icon: {
    marginRight: '5pt'
  },
  inputBox: {
    display: 'flex',
    alignItems: 'center',
    border: '1px solid lightgray',
    borderRadius: '4px',
    padding: '3pt 0 0 2pt',
    margin: '5pt 15pt 5pt 60pt',
    width: '40%'
  },
  input: {
    paddingLeft: '2pt',
  },
  addButton: {
    margin: '5pt 155pt 5pt auto',
    paddingLeft: '4pt',
  },
}));

function TabPanel(props) {
  const {children, value, index, ...other} = props;

  return (
    <Typography
      component="div"
      role="tabpanel"
      className={props.className}
      hidden={value !== index}
      id={`vertical-tabpanel-${index}`}
      aria-labelledby={`vertical-tab-${index}`}
      {...other}
    >
      {value === index && <Box className={props.flex} p={0}>{children}</Box>}
    </Typography>
  );
}

function a11yProps(index) {
  return {
    id: `vertical-tab-${index}`,
    'aria-controls': `vertical-tabpanel-${index}`,
  };
}

function SearchInput(props) {
  const styles = useStyles();

  return (
    <div className={styles.inputBox}>
      <SearchIcon/>
      <InputBase fullWidth
                 onChange={props.onChange}
                 placeholder={props.placeholder}
                 className={styles.input}
                 inputProps={{'aria-label': 'search'}}
      />
    </div>
  )
}

function SubjectPage(props) {
  const styles = useStyles();
  const classes = useStyles();
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

  return (
    <>
      <MaterialsPathHeader university={props.university} subject={props.subject}/>
      <div className={classes.wrapper}>
        <div className={classes.root}>
          <Tabs
            value={value}
            onChange={handleChange}
            classes={{indicator: classes.indicator}}
            indicatorColor="primary"
            textColor="primary"
            aria-label="full width tabs example"
          >
            <Tab classes={{root: styles.tab}}
                 label={<div className={styles.label}><FolderRoundedIcon className={styles.icon}/>Файлы</div>}
                 {...a11yProps(0)} />
            <Tab classes={{root: styles.tab}}
                 label={<div className={styles.label}><GroupRoundedIcon className={styles.icon}/>Комнаты</div>}
                 {...a11yProps(1)} />
            <Tab classes={{root: styles.tab}}
                 label={<div className={styles.label}><FormatListNumberedRoundedIcon className={styles.icon}/>Тесты
                 </div>}
                 {...a11yProps(2)} />
          </Tabs>
          <TabPanel className={classes.highPanel} flex={classes.flex}
                    value={value} index={0}>
            <SearchInput onChange={props.onNameChange} placeholder={'Название...'}/>
            <Button href={`/add_material/${props.subject.id}`}
                    className={classes.addButton} variant="contained"
                    color="primary">
              <AddIcon/> Добавить материал
            </Button>
          </TabPanel>
          <TabPanel className={classes.highPanel} flex={classes.flex}
                    value={value} index={1}>
            <SearchInput onChange={props.onRoomNameChange} placeholder={'Название...'}/>
            <Button onClick={handleClickOpen}
                    className={classes.addButton} variant="contained"
                    color="primary">
              <AddIcon/> Создать комнату
            </Button>
          </TabPanel>
        </div>
        <hr className={classes.hr}/>
        <div className={classes.tabPanelWrapper}>
          <div className={classes.left}>
            <TabPanel className={classes.tabPanel} value={value} index={0}>
              <TypeFilter onSemesterChange={props.onTypeChange}/>
            </TabPanel>
            <TabPanel className={classes.tabPanel} value={value} index={1}>
              <TypeFilter onSemesterChange={props.onRoomTypeChange}/>
            </TabPanel>
          </div>
          <TabPanel className={classes.tabPanel} value={value} index={0}>
            <Container maxWidth="lg">
              <MaterialTab id={props.id}
                           data={props.data}
                           pageCount={props.pageCount}
                           currPage={props.currPage}
                           onChange={props.onChange}
                           onNameChange={props.onNameChange}
                           onTypeChange={props.onTypeChange}/>
            </Container>
          </TabPanel>
          <TabPanel className={classes.tabPanel} value={value} index={1}>
            <RoomsTab id={props.id}
                      onNameChange={props.onNameChange}
                      onTypeChange={props.onTypeChange}/>
          </TabPanel>
          <TabPanel value={value} index={2}>
            Item Three
          </TabPanel>
        </div>
      </div>
      <AddDialog open={open}
                 universities={props.universities}
                 subjects={props.subjects}
                 university={props.university.id}
                 semester={props.subject.semester}
                 subject={props.subject.id}
                 handleClose={handleClose}/>
    </>
  );
}

class SubjectPageController extends React.Component {
  constructor(props) {
    super(props);

    this.currPage = 1;
    this.pageCount = 1;
    this.universityService = new UniversitiesService();
    this.subjectsService = new SubjectsService();
    this.materialService = new MaterialService();
  }

  university = {name: ''};
  subject = {name: ''};
  materials = [];
  universities = [];
  subjects = [];

  componentDidMount() {
    const page = 1; // TODO get from url
    this.getSubject();
    this.getMaterials(page);
    this.getUniversitiesList();
  }

  getSubject = () => {
    const id = this.props.id;
    this.subjectsService.getById(id).then((result) => {
        runInAction(() => {
          this.subject = result;
        });
        this.getUniversity();
        this.getSubjectsList(this.subject.university_id);
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

  getMaterials = (page, name, type) => {
    const id = this.props.id;
    this.currPage = page;
    console.log(name, type);
    this.materialService.getPage(id, page, name, type).then((result) => {
        runInAction(() => {
          this.pageCount = result.pages;
          this.materials = result.payload ? result.payload : [];
        });
      },
      (error) => {
        console.log(error);
      });
  };

  onPageClick = (event, page) => {
    this.getMaterials(page);
  };

  onTypeChange = (event) => {
    this.type = event.target.value;
    this.getMaterials(1, this.name, this.type);
  };

  onNameChange = (event) => {
    this.name = event.target.value;
    this.getMaterials(1, this.name, this.type);
  };

  onRoomNameChange = (event) => {
    this.roomName = event.target.value;
  };

  onRoomTypeChange = (event) => {
    this.roomType = event.target.value;
  };

  getUniversitiesList = () => {
    const universitiesService = new UniversitiesService();
    universitiesService.getAll().then((result) => {
      runInAction(() => {this.universities = result.payload;});
    });
  };

  getSubjectsList = (university_id) => {
    const subjectsService = new SubjectsService();
    subjectsService.getAll(university_id).then((result) => {
      runInAction(() => {this.subjects = result.payload;});
      console.log(this.subjects);
    });
  };

  render() {
    return (
      <SubjectPage id={this.subject.id}
                   university={this.university}
                   subject={this.subject}
                   universities={this.universities}
                   subjects={this.subjects}
                   data={this.materials}
                   currPage={this.currPage}
                   pageCount={this.pageCount}
                   onTypeChange={this.onTypeChange}
                   onNameChange={this.onNameChange}
                   onRoomNameChange={this.onRoomNameChange}
                   onRoomTypeChange={this.onRoomTypeChange}
                   onChange={this.onPageClick}/>
    )
  }
}

decorate(SubjectPageController, {
  pageCount: observable,
  subject: observable,
  university: observable,
  materials: observable,
  rooms: observable,
  universities: observable,
  subjects: observable,
});

export default observer(SubjectPageController);
