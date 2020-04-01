import React from "react";
import {makeStyles} from "@material-ui/core/styles";
import FormControl from "@material-ui/core/FormControl";
import FormLabel from "@material-ui/core/FormLabel";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Radio from "@material-ui/core/Radio";

const useStyles = makeStyles(() => ({
  form: {
    marginTop: '10pt',
    marginLeft: '55pt',
    marginBottom: '5pt',
  },
  mainLabel: {
    color: 'black',
    fontWeight: 'bold',
    fontSize: 'x-large',
  },
  label: {
    fontSize: 'large',
  }
}));

function SemesterFilter(props) {
  const classes = useStyles();
  const [value, setValue] = React.useState('any');

  const handleChange = (event) => {
    setValue(event.target.value);
    props.onSemesterChange(event);
  };

  return (
    <FormControl className={classes.form} component="fieldset">
      <FormLabel className={classes.mainLabel} component="legend">Семестр</FormLabel>
      <RadioGroup aria-label="semester" name="semester" value={value} onChange={handleChange}>
        <FormControlLabel value={'any'} control={<Radio color={'primary'}/>} label="Любой" />
        <FormLabel className={classes.label}>Бакалавриат</FormLabel>
        <FormControlLabel value={'1'} control={<Radio color={'primary'}/>} label="Первый" />
        <FormControlLabel value={'2'} control={<Radio color={'primary'}/>} label="Второй" />
        <FormControlLabel value={'3'} control={<Radio color={'primary'}/>} label="Третий" />
        <FormControlLabel value={'4'} control={<Radio color={'primary'}/>} label="Четвертый" />
        <FormControlLabel value={'5'} control={<Radio color={'primary'}/>} label="Пятый" />
        <FormControlLabel value={'6'} control={<Radio color={'primary'}/>} label="Шестой" />
        <FormControlLabel value={'7'} control={<Radio color={'primary'}/>} label="Седьмой" />
        <FormControlLabel value={'8'} control={<Radio color={'primary'}/>} label="Восьмой" />
        <FormLabel className={classes.label}>Специалитет</FormLabel>
        <FormControlLabel value={'9'} control={<Radio color={'primary'}/>} label="Девятый" />
        <FormControlLabel value={'10'} control={<Radio color={'primary'}/>} label="Десятый" />
        <FormLabel className={classes.label}>Магистратура</FormLabel>
        <FormControlLabel value={'1М'} control={<Radio color={'primary'}/>} label="Первый" />
        <FormControlLabel value={'2М'} control={<Radio color={'primary'}/>} label="Второй" />
        <FormControlLabel value={'3М'} control={<Radio color={'primary'}/>} label="Третий" />
        <FormControlLabel value={'4М'} control={<Radio color={'primary'}/>} label="Четвертый" />
      </RadioGroup>
    </FormControl>
  );
}

export default SemesterFilter;
