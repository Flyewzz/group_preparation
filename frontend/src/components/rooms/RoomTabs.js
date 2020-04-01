import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import FolderRoundedIcon from '@material-ui/icons/FolderRounded';
import GroupRoundedIcon from '@material-ui/icons/GroupRounded';
import PanTool from '@material-ui/icons/PanTool'
import AddIcon from '@material-ui/icons/Add'
import FormatListNumberedRoundedIcon from '@material-ui/icons/FormatListNumberedRounded';
import Button from "@material-ui/core/Button";

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

function RoomTabs(props) {
  const styles = useStyles();
  const classes = useStyles();
  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  return (
    <>
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
                 label={<div className={styles.label}><PanTool className={styles.icon}/>Доска</div>}
                 {...a11yProps(1)} />
            <Tab classes={{root: styles.tab}}
                 label={<div className={styles.label}><FormatListNumberedRoundedIcon className={styles.icon}/>Тесты
                 </div>}
                 {...a11yProps(2)} />
          </Tabs>
        </div>
        <hr className={classes.hr}/>
    </>
  );
}

export default RoomTabs;