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
    paddingBottom: '10px',
    paddingTop: '10px',
  },
  name: {
    fontSize: 'large',
    fontWeight: 'bold',
  },
  text: {
    fontSize: 'large',
  },
}));

function RoomsListItem(props) {
  const styles = useStyles();

  return (
    <TableRow className={styles.listItem}>
      <TableCell className={styles.cell}>
        <Link component={RouterLink} to={'/room/' + props.room.id} className={styles.listItem} underline="none"
              color="inherit">
          <ListItemText primary={props.room.name}
                        primaryTypographyProps={{className: styles.name}}/>
        </Link>
      </TableCell>
      <TableCell align={'center'} className={styles.cell}>
        <Link component={RouterLink} to={'/room/' + props.room.room_id} className={styles.listItem} underline="none"
              color="inherit">
          <ListItemText primaryTypographyProps={{className: styles.text}}
                        primary={props.room.type}/>
        </Link>
      </TableCell>
      <TableCell align={'center'} className={styles.cell}>
        <Link component={RouterLink} to={'/room/' + props.room.room_id} className={styles.listItem} underline="none"
              color="inherit">
          <ListItemText primaryTypographyProps={{className: styles.text}}
                        primary={props.room.author}/>
        </Link>
      </TableCell>
      <TableCell align="right" className={styles.cell}>
        <Link component={RouterLink} to={'/room/' + props.room.room_id} className={styles.listItem} underline="none"
              color="inherit">
          <ListItemText primary={'1'}
                        primaryTypographyProps={{className: styles.text}}/>
        </Link>
      </TableCell>
    </TableRow>
  );
}

export default RoomsListItem;
