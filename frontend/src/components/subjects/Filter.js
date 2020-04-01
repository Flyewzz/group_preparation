import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import SearchIcon from "@material-ui/icons/Search"
import InputBase from "@material-ui/core/InputBase";

const useStyles = makeStyles(() => ({
  filter: {
    display: 'flex',
    alignItems: 'center',
    marginTop: '6pt'
  },
  inputBox: {
    fontSize: '30px',
    display: 'flex',
    alignItems: 'center',
    border: '1px solid lightgray',
    borderRadius: '4px',
    padding: '3pt 0 0 2pt',
    width: '100%'
  },
  input: {
    fontSize: 'large',
    paddingLeft: '2pt',
  },
}));

function Filter(props) {
  const styles = useStyles();

  return (
    <div className={styles.filter}>
      <div className={styles.inputBox}>
        <SearchIcon fontSize={'inherit'}/>
        <InputBase fullWidth
                   onChange={props.onNameChange}
                   placeholder={'Название...'}
                   className={styles.input}
                   inputProps={{'aria-label': 'search'}}
        />
      </div>
    </div>
  );
}

export default Filter;
