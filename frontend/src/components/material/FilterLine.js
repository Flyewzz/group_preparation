import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import SearchInput from "../common/SearchInput";
import FormControl from "@material-ui/core/FormControl";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import Button from "@material-ui/core/Button";
import InputLabel from "@material-ui/core/InputLabel";
import withStyles from "@material-ui/core/styles/withStyles";

const useStyles = makeStyles(() => ({
  wrapper: {
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    margin: '9pt 15pt 0 15pt'
  },
  formControl: {
    minWidth: 180,
    marginRight: '15pt',
  },
  select: {
    padding: '10px 8px 10px 14px',
  },
  label: {
    transform: 'translate(14px, 12px) scale(1)'
  },
  button: {
    margin: 'auto 7px auto 15px'
  },
}));

const AddButton = withStyles(theme => ({
  label: {
    margin: 'auto 7px auto 7px',
    minWidth: '150px'
  },
}))(Button);

function FilterLine(props) {
  const styles = useStyles();
  const [type, setType] = React.useState('');
  const handleChange = event => {
    setType(event.target.value);
  };

  return (
    <div className={styles.wrapper}>
      <SearchInput placeholder={'Название...'}/>
      <FormControl variant="outlined" className={styles.formControl}>
        <InputLabel className={styles.label}
                    id="demo-simple-select-outlined-label">
          Тип мероприятия
        </InputLabel>
        <Select
          classes={{select: styles.select}}
          labelId="demo-simple-select-outlined-label"
          id="demo-simple-select-outlined"
          value={type}
          onChange={handleChange}
          label="Тип мероприятия"
        >
          <MenuItem value="">
            <em>Любой</em>
          </MenuItem>
          <MenuItem value={1}>РК</MenuItem>
          <MenuItem value={2}>ЛР</MenuItem>
          <MenuItem value={3}>ДЗ</MenuItem>
          <MenuItem value={4}>КР</MenuItem>
          <MenuItem value={5}>Экз</MenuItem>
          <MenuItem value={6}>Зачет</MenuItem>
          <MenuItem value={7}>Лекции</MenuItem>
          <MenuItem value={8}>Семинар</MenuItem>
          <MenuItem value={9}>Методы</MenuItem>
        </Select>
      </FormControl>
      <AddButton disableElevation variant="contained" color="primary">
        Добавить
      </AddButton>
    </div>
  );
}

export default FilterLine;
