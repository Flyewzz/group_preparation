import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import SearchInput from "../common/SearchInput";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";

const useStyles = makeStyles(() => ({
  filter: {
    display: 'flex',
    alignItems: 'center',
    margin: '6pt 11pt 0 11pt'
  },
  formControl: {
    minWidth: 120,
  },
  select: {
    padding: '10px 8px 10px 14px',
  },
  label: {
    transform: 'translate(14px, 12px) scale(1)'
  },
}));

function Filter(props) {
  const styles = useStyles();
  const [semester, setSemester] = React.useState('');
  const handleChange = event => {
    setSemester(event.target.value);
    props.onSemesterChange(event);
  };

  return (
    <div className={styles.filter}>
      <SearchInput onChange={props.onNameChange} placeholder={'Название...'}/>
      <FormControl variant="outlined" className={styles.formControl}>
        <InputLabel className={styles.label}
                    id="demo-simple-select-outlined-label">
          Семестр
        </InputLabel>
        <Select
          classes={{select: styles.select}}
          labelId="demo-simple-select-outlined-label"
          id="demo-simple-select-outlined"
          value={semester}
          onChange={handleChange}
          label="Семестр"
        >
          <MenuItem value="">
            <em>Любой</em>
          </MenuItem>
          <MenuItem value={1}>1ый</MenuItem>
          <MenuItem value={2}>2ой</MenuItem>
          <MenuItem value={3}>3ий</MenuItem>
          <MenuItem value={4}>4ый</MenuItem>
          <MenuItem value={5}>5ый</MenuItem>
          <MenuItem value={6}>6ой</MenuItem>
          <MenuItem value={7}>7ой</MenuItem>
          <MenuItem value={8}>8ой</MenuItem>
          <MenuItem value={9}>9ый</MenuItem>
          <MenuItem value={10}>10ый</MenuItem>
        </Select>
      </FormControl>
    </div>
  );
}

export default Filter;
