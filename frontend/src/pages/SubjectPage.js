import React from "react";
import Material from "../components/material/Material";
import Container from "@material-ui/core/Container";
import {makeStyles} from "@material-ui/core/styles";
import Pagination from "@material-ui/lab/Pagination";
import TableHeader from "../components/material/TableHeader"
import FilterLine from "../components/material/FilterLine";
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import FolderRoundedIcon from '@material-ui/icons/FolderRounded';
import GroupRoundedIcon from '@material-ui/icons/GroupRounded';
import FormatListNumberedRoundedIcon from '@material-ui/icons/FormatListNumberedRounded';


const data = [
  {
    name: 'РК №1',
    department: 'ИУ7',
    year: '2020',
    type: 'РК',
    author: 'username',
    rating: 23,
  },
  {
    name: 'ДЗ №1',
    department: 'ИУ7',
    year: '2020',
    type: 'ДЗ',
    author: 'username',
    rating: 20,
  },
  {
    name: 'ЛР №1',
    department: 'ИУ7',
    year: '2020',
    type: 'ЛР',
    author: 'username',
    rating: -13,
  },
  {
    name: 'РК №2',
    year: '2020',
    type: 'РК',
    author: 'username',
    rating: 23,
  },
  {
    name: 'ДЗ №2',
    department: 'ИУ7',
    year: '2020',
    type: 'ДЗ',
    author: 'username',
    rating: 0,
  },
  {
    name: 'РК №1',
    department: 'ИУ7',
    year: '2020',
    type: 'РК',
    author: 'username',
    rating: 23,
  },
  {
    name: 'ДЗ №1',
    department: 'ИУ7',
    year: '2020',
    type: 'ДЗ',
    author: 'username',
    rating: 20,
  },
  {
    name: 'ЛР №1',
    department: 'ИУ7',
    year: '2020',
    type: 'ЛР',
    author: 'username',
    rating: -13,
  },
];

const useStyles = makeStyles((theme) => ({
  wrapper: {
    width: '95%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '0',
    border: '1px solid #535455',
    borderRadius: '0px 6px 6px 6px',
    padding: '10pt',
  },
  tabsWrapper: {
    width: '95%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '20pt auto',
    borderRadius: '6px',
    padding: '10pt',
  },
  table: {
    width: '100%',
    marginBottom: '10pt',
  },
  header: {
    fontWeight: 'bold',
    marginBottom: '5pt',
  },
  list: {
    width: '95%',
  },
  name: {
    fontWeight: 'bold',
    fontSize: 'xx-large'
  },
  root: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.paper,
    display: 'flex',
    height: 224,
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
}));


function TabPanel(props) {
  const {children, value, index, ...other} = props;

  return (
    <Typography
      component="div"
      role="tabpanel"
      hidden={value !== index}
      id={`vertical-tabpanel-${index}`}
      aria-labelledby={`vertical-tab-${index}`}
      {...other}
    >
      {value === index && <Box p={0}>{children}</Box>}
    </Typography>
  );
}

function a11yProps(index) {
  return {
    id: `vertical-tab-${index}`,
    'aria-controls': `vertical-tabpanel-${index}`,
  };
}

const getSubjectName = () => {
  return 'Физика';
};

function SubjectPage(props) {
  const styles = useStyles();
  const classes = useStyles();
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  return (
    <Container maxWidth="md" className={styles.tabsWrapper}>
      <div className={classes.root}>
        <Tabs orientation="vertical"
              centered
              value={value}
              onChange={handleChange}
              aria-label="Vertical tabs example"
              className={classes.tabs}
              indicatorColor="primary"
        >
          <Tab classes={{root: styles.tab}} icon={<FolderRoundedIcon/>} label="Файлы" {...a11yProps(0)} />
          <Tab classes={{root: styles.tab}} icon={<GroupRoundedIcon/>} label="Комнаты" {...a11yProps(1)} />
          <Tab classes={{root: styles.tab}} icon={<FormatListNumberedRoundedIcon/>} label="Тесты" {...a11yProps(2)} />
        </Tabs>
        <TabPanel value={value} index={0}>
          <Container maxWidth="sm" className={styles.wrapper}>
            <div className={styles.name}>
              {getSubjectName()}
            </div>
            <FilterLine/>
            <table className={styles.table}>
              <TableHeader/>
              {data.map((value) =>
                <Material material={value}/>
              )}
            </table>
            <Pagination count={10} shape="rounded" color="primary"/>
          </Container>
        </TabPanel>
        <TabPanel value={value} index={1}>
          <Container maxWidth="sm" className={styles.wrapper}>
            Item Two
          </Container>
        </TabPanel>
        <TabPanel value={value} index={2}>
          <Container maxWidth="sm" className={styles.wrapper}>
            Item Three
          </Container>
        </TabPanel>
      </div>
    </Container>
  );
}

export default SubjectPage;