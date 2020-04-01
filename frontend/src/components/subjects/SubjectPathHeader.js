import React from "react"
import {makeStyles} from "@material-ui/core/styles";
import Link from "@material-ui/core/Link";
import {Link as RouterLink} from "react-router-dom";

const useStyles = makeStyles(() => ({
  wrapper: {
    fontSize: 'large',
    margin: '15pt 15pt 15pt 55pt',
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

function SubjectPathHeader(props) {
  const classes = useStyles();

  return (
    <div className={classes.wrapper}>
      <Link component={RouterLink} className={classes.enabled} to={'/'}>Университеты</Link>
      <Link className={classes.disabled} disabled>
        {' > ' + props.university.name + ': Предметы'}
      </Link>
    </div>
  );
}

export default SubjectPathHeader;
