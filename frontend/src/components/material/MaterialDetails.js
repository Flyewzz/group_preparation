import React from "react";
import {Link as RouterLink} from "react-router-dom";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import AccountBalanceRoundedIcon from '@material-ui/icons/AccountBalanceRounded';
import AssignmentRoundedIcon from '@material-ui/icons/AssignmentRounded';
import AssessmentRoundedIcon from '@material-ui/icons/AssessmentRounded';
import PersonRoundedIcon from '@material-ui/icons/PersonRounded';
import SchoolRoundedIcon from '@material-ui/icons/SchoolRounded';
import InsertInvitationRoundedIcon from '@material-ui/icons/InsertInvitationRounded';
import FolderRoundedIcon from '@material-ui/icons/FolderRounded';
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  table: {
    paddingTop: '10pt',
    width: '100%',
  },
  listItem: {
    width: '100%',
  },
  icon: {
    padding: '4px 0 0 0'
  },
  name: {
    fontSize: 'large',
    fontWeight: 'bold',
    paddingLeft: 0,
    paddingBottom: '4px',
    paddingTop: '4px',
  },
  cell: {
    fontSize: 'large',
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

function Line(props) {
  const styles = useStyles();

  return (
    <TableRow className={styles.listItem}>
      <TableCell align={'center'} className={styles.icon}>
        {props.icon}
      </TableCell>
      <TableCell align={'left'} className={styles.name}>
        {props.name}
      </TableCell>
      <TableCell align={'right'} className={styles.cell}>
        {props.value}
      </TableCell>
    </TableRow>
  );
}

function MaterialDetails(props) {
  const styles = useStyles();
  const date = new Date(props.material.date.replace(' ', 'T')).toLocaleDateString();

  return (
    <table className={styles.table}>
      <tbody>
      <Line name={'ВУЗ'}
            value={'МГТУ им. Н.Э.Баумана'} // props.material.university}
            icon={<AccountBalanceRoundedIcon/>}/>
      <Line name={'Предмет'}
            value={'Дискретная математика'} // props.material.subject}
            icon={<AssignmentRoundedIcon/>}/>
      <Line name={'Семестр'}
            value={'4ый'} // props.material.semester}
            icon={<SchoolRoundedIcon/>}/>
      <Line name={'Тип'}
            value={props.material.type}
            icon={<FolderRoundedIcon/>}/>
      <Line name={'Дата'}
            value={date}
            icon={<InsertInvitationRoundedIcon/>}/>
      <Line name={'Автор'}
            value={props.material.user_email}
            icon={<PersonRoundedIcon/>}/>
      <Line name={'Рейтинг'}
            value={0} // props.material.rating
            icon={<AssessmentRoundedIcon/>}/>
      </tbody>
    </table>
  );
}

export default MaterialDetails;
