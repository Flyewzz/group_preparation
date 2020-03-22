import React from "react";
import Typography from "@material-ui/core/Typography"
import Divider from "@material-ui/core/Divider";
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  wrapper: {
    display: 'flex',
    margin: '15pt 11pt 0 11pt'
  },
  first: {
    fontWeight: 'bold',
    marginRight: 'auto',
  },
}));

function Header(props) {
  const styles = useStyles();

  return (
    <>
      <div className={styles.wrapper}>
        <Typography className={styles.first} variant="h5">
          Предмет
        </Typography>
        <Typography variant="h6">
          Файлов
        </Typography>
      </div>
      <Divider variant="middle"/>
    </>
  );
}

export default Header;
