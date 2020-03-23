import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";

const useStyles = makeStyles(() => ({
  bold: {
    fontWeight: 'bold',
    fontSize: 'large',
  },
}));

function TableHeader(props) {
  const styles = useStyles();

  return (
    <TableHead>
      <TableRow>
        <TableCell className={styles.bold}>Название</TableCell>
        <TableCell align={'center'} className={styles.bold}>Тип</TableCell>
        <TableCell align={'center'} className={styles.bold}>Автор</TableCell>
        <TableCell className={styles.bold} align="right">Рейтинг</TableCell>
      </TableRow>
    </TableHead>
  );
}

export default TableHeader;
