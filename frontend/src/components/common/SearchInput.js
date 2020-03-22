import React from "react";
import SearchIcon from "@material-ui/icons/Search"
import InputBase from "@material-ui/core/InputBase"
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  inputBox: {
    display: 'flex',
    alignItems: 'center',
    border: '1px solid lightgray',
    borderRadius: '4px',
    padding: '3pt 0 0 2pt',
    marginRight: '15pt',
    width: '100%'
  },
  input: {
    paddingLeft: '2pt',
  },
}));

function SearchInput(props) {
  const styles = useStyles();

  return (
    <div className={styles.inputBox}>
      <SearchIcon/>
      <InputBase fullWidth
                 placeholder={props.placeholder}
                 className={styles.input}
                 inputProps={{'aria-label': 'search'}}
      />
    </div>
  )
}

export default SearchInput;
