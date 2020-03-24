import React from "react";
import ListItemText from "@material-ui/core/ListItemText";
import Link from "@material-ui/core/Link";
import {Link as RouterLink} from "react-router-dom";
import {makeStyles} from "@material-ui/core/styles";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";

const useStyles = makeStyles(() => ({
  listItem: {
    width: '100%',
  },
  cell: {
    paddingBottom: '4px',
    paddingTop: '4px',
  },
  positive: {
    fontWeight: 'bold',
    color: '#216416',
  },
  negative: {
    fontWeight: 'bold',
    color: '#A01919',
  },
  zero: {
    fontWeight: 'bold',
    color: '#6B6B6B',
  }
}));

const getRatingStyle = (styles, rating) => {
  if (rating > 0) {
    return styles.positive;
  } else if (rating < 0) {
    return styles.negative;
  } else {
    return styles.zero;
  }
};

const addSign = (rating) => {
  if (rating > 0) {
    return '+' + rating;
  }
  return rating;
};

function Material(props) {
  const styles = useStyles();
  const department = props.material.department ? props.material.department + ', ' : '';
  const secondaryText = department + props.material.year;

  return (
    <TableRow className={styles.listItem}>
      <TableCell className={styles.cell}>
        <Link component={RouterLink} to={'/material/' + props.material.id} className={styles.listItem} underline="none" color="inherit">
          <ListItemText primary={props.material.name}
                        secondary={secondaryText}
                        className={styles.first}/>
        </Link>
      </TableCell>
      <TableCell align={'center'} className={styles.cell}>
        <Link component={RouterLink} to={'/material/' + props.material.id} className={styles.listItem} underline="none" color="inherit">
          <ListItemText primary={props.material.type}/>
        </Link>
      </TableCell>
      <TableCell align={'center'} className={styles.cell}>
        <Link component={RouterLink} to={'/material/' + props.material.id} className={styles.listItem} underline="none" color="inherit">
          <ListItemText className={styles.username}
                        primary={props.material.author}/>
        </Link>
      </TableCell>
      <TableCell align="right" className={styles.cell}>
        <Link component={RouterLink} to={'/material/' + props.material.id} className={styles.listItem} underline="none" color="inherit">
          <ListItemText primary={addSign(props.material.rating)}
                        primaryTypographyProps={
                          {className: getRatingStyle(styles, props.material.rating)}
                        }/>
        </Link>
      </TableCell>
    </TableRow>
  );
}

export default Material;
