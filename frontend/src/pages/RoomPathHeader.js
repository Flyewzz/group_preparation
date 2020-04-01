import React from "react"
import {makeStyles} from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import {Link as RouterLink} from "react-router-dom";

const useStyles = makeStyles(() => ({
  wrapper: {
    fontSize: 'large',
    margin: '15pt 15pt 18pt 55pt',
  },
  enabled: {
    color: 'black',
    cursor: 'pointer',
  },
  disabled: {
    '&:hover': {
      textDecoration: 'none',
    },
    color: 'gray'
  },
}));

function RoomPathHeader(props) {
  const classes = useStyles();

  return (
    <div className={classes.wrapper}>
      <Link component={RouterLink}
            className={classes.enabled} to={'/'}>
        Университеты
      </Link>
      <Link component={RouterLink}
            className={classes.enabled}
            to={`/university/${props.university.id}`}>
        {' > ' + props.university.name}
      </Link>
      <Link component={RouterLink}
            className={classes.disabled}
            to={`/subject/${props.subject.id}`}>
        {' > ' + props.subject.name}
      </Link>
      <Link component={RouterLink}
            disabled
            className={classes.disabled}>
        {' > ' + props.roomName}
      </Link>
    </div>
  );
}

export default RoomPathHeader;
