import Container from "@material-ui/core/Container"
import Typography from "@material-ui/core/Typography";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import Divider from "@material-ui/core/Divider"
import Pagination from "@material-ui/lab/Pagination"
import {makeStyles} from "@material-ui/core/styles";
import React from "react";

const useStyles = makeStyles(() => ({
  wrapper: {
    width: '95%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '20pt auto',
    padding: '10pt',
  },
  header: {
    fontWeight: 'bold',
    marginBottom: '5pt',
  },
  list: {
    width: '95%',
  },
}));

function formList(items) {
  const result = [];
  items.forEach((value, key) => {
    result.push(<ListItem button key={key}>
      {value}
    </ListItem>);

    if (key !== items.length - 1) {
      result.push(
        <Divider key={items.length + key}
                 variant="middle" component="li"/>
      );
    }
  });

  return result;
}

function ListContainer(props) {
  const styles = useStyles();

  return (
    <Container maxWidth="lg" className={styles.wrapper}>
      <Typography className={styles.header} variant="h4">
        {props.title}
      </Typography>
      <List className={styles.list} subheader={props.subheader}>
        {formList(props.items)}
      </List>
      <Pagination count={props.pageCount}
                  page={props.currPage}
                  onChange={props.onChange}
                  size="large"
                  shape="rounded"
                  color="primary"/>
    </Container>
  );
}

export default ListContainer;
