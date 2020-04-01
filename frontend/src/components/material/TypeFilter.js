import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import FormControl from "@material-ui/core/FormControl";
import FormLabel from "@material-ui/core/FormLabel";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Radio from "@material-ui/core/Radio";

const useStyles = makeStyles(() => ({
  form: {
    marginTop: '25pt',
    marginBottom: '5pt',
  },
  mainLabel: {
    width: 'max-content',
    color: 'black',
    fontWeight: 'bold',
    fontSize: 'x-large',
  },
  label: {
    width: 'max-content',
  }
}));

function TypeFilter(props) {
  const classes = useStyles();
  const [value, setValue] = React.useState('');

  const handleChange = (event) => {
    setValue(event.target.value);
    props.onSemesterChange(event);
  };

  return (
    <FormControl className={classes.form} component="fieldset">
      <FormLabel className={classes.mainLabel} component="legend">Тип мероприятия</FormLabel>
      <RadioGroup aria-label="type" name="type" value={value} onChange={handleChange}>
        <FormControlLabel value={''} control={<Radio color={'primary'}/>} label="Любой" />
        <FormControlLabel value={'1'}
                          className={classes.label}
                          control={<Radio color={'primary'}/>}
                          label="Рубежный контроль" />
        <FormControlLabel value={'2'}
                          className={classes.label}
                          control={<Radio color={'primary'}/>}
                          label="Лабораторная работа" />
        <FormControlLabel value={'3'}
                          control={<Radio color={'primary'}/>}
                          label="Домашнее задание" />
        <FormControlLabel value={'4'}
                          control={<Radio color={'primary'}/>}
                          label="Контрольная работа" />
        <FormControlLabel value={'5'}
                          control={<Radio color={'primary'}/>} label="Лекции" />
        <FormControlLabel value={'6'} control={<Radio color={'primary'}/>} label="Семинары" />
        <FormControlLabel value={'7'} control={<Radio color={'primary'}/>} label="Зачет" />
        <FormControlLabel value={'8'} control={<Radio color={'primary'}/>} label="Экзамен" />
        <FormControlLabel value={'9'} control={<Radio color={'primary'}/>} label="Методички" />
      </RadioGroup>
    </FormControl>
  );
}

export default TypeFilter;
