import React from "react";
import TableHeader from "./TableHeader";
import Pagination from "@material-ui/lab/Pagination";
import {makeStyles} from "@material-ui/core/styles";
import RoomsListItem from "./RoomsListItem";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles(() => ({
  wrapper: {
    width: '100%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '10pt 0',
  },
  table: {
    width: '100%',
    marginBottom: '10pt',
  },
  list: {
    width: '95%',
  },
}));

function RoomsList(props) {
  const styles = useStyles();

  return (
    <Container maxWidth="lg">
      <div className={styles.wrapper}>
        <table className={styles.table}>
          <TableHeader/>
          {props.data.map((value) =>
            <RoomsListItem room={value}/>
          )}
        </table>
        {props.pageCount > 1 && <Pagination count={props.pageCount}
                    page={props.currPage}
                    onChange={props.onChange}
                    shape="rounded"
                    color="primary"/>}
      </div>
    </Container>
  );
}

export default RoomsList;