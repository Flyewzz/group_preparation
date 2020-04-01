import React from "react";
import TableHeader from "./TableHeader";
import Material from "./Material";
import Pagination from "@material-ui/lab/Pagination";
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
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

function MaterialTab(props) {
  const styles = useStyles();

  return (
    <div className={styles.wrapper}>
      <table className={styles.table}>
        <TableHeader/>
        {props.data.map((value) =>
          <Material material={value}/>
        )}
      </table>
      <Pagination count={props.pageCount}
                  page={props.currPage}
                  onChange={props.onChange}
                  shape="rounded"
                  color="primary"/>
    </div>
  );
}

export default MaterialTab;