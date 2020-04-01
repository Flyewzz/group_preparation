import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import Grid from "@material-ui/core/Grid";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(3),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  wrapper: {
    width: '95%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    margin: '20pt auto',
    border: '1px solid #535455',
    borderRadius: '6px',
    padding: '10pt',
  },
  formControl: {
    minWidth: '100%',
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
}));

export default function AddDialog(props) {
  const classes = useStyles();
  const [type, setType] = React.useState('');
  const handleTypeChange = event => {
    setType(event.target.value);
  };
  const [semester, setSemester] = React.useState('');
  if (props.semester && semester !== props.semester) {
    setSemester(props.semester);
  }
  const handleSemesterChange = event => {
    setSemester(event.target.value);
  };
  const [university, setUniversity] = React.useState('');
  const handleUniversityChange = event => {
    setUniversity(event.target.value);
  };
  if (props.university && university !== props.university) {
    handleUniversityChange({target: {value: props.university}});
  }
  const [subject, setSubject] = React.useState('');
  const handleSubjectChange = event => {
    setSubject(event.target.value);
  };
  if (props.subject && subject !== props.subject) {
    setSubject(props.subject);
    handleSubjectChange({target: {value: props.subject}});
  }

  return (
    <div>
      <Dialog open={props.open} onClose={props.handleClose} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Создать комнату</DialogTitle>
        <DialogContent>
          <form className={classes.form} noValidate>
            <Grid container spacing={2}>
              <Grid container spacing={2}>
                <Grid item xs={12} sm={9}>
                  <TextField
                    size={'small'}
                    autoComplete="mname"
                    name="materialName"
                    variant="outlined"
                    required
                    fullWidth
                    id="materialName"
                    label="Название"
                    autoFocus
                  />
                </Grid>
                <Grid item xs={12} sm={3}>
                  <FormControl size={'small'} required variant="outlined" className={classes.formControl}>
                    <InputLabel htmlFor="university">Тип</InputLabel>
                    <Select
                      value={type}
                      onChange={handleTypeChange}
                      label="Тип"
                      inputProps={{
                        name: 'type_id',
                        id: 'university-select',
                      }}
                    >
                      <MenuItem value="">
                        <em>Любой</em>
                      </MenuItem>
                      <MenuItem value={1}>РК</MenuItem>
                      <MenuItem value={2}>ЛР</MenuItem>
                      <MenuItem value={3}>ДЗ</MenuItem>
                      <MenuItem value={4}>КР</MenuItem>
                      <MenuItem value={5}>Зачет</MenuItem>
                      <MenuItem value={6}>Экзамен</MenuItem>
                      <MenuItem value={7}>Лекции</MenuItem>
                      <MenuItem value={8}>Семинары</MenuItem>
                      <MenuItem value={9}>Методички</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} sm={9}>
                  <FormControl size={'small'} required variant="outlined" className={classes.formControl}>
                    <InputLabel htmlFor="university">Университет</InputLabel>
                    <Select
                      value={university}
                      onChange={handleUniversityChange}
                      label="Университет *"
                      inputProps={{
                        name: 'university',
                        id: 'university-select',
                      }}
                    >
                      <MenuItem value="">
                        <em>Любой</em>
                      </MenuItem>
                      {props.universities && props.universities.map((item) => (
                        <MenuItem value={item.id}>{item.name}</MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} sm={3}>
                  <FormControl disabled={university === ''} size={'small'} variant="outlined" className={classes.formControl}>
                    <InputLabel htmlFor="outlined-age-native-simple">Семестр</InputLabel>
                    <Select
                      value={semester}
                      onChange={handleSemesterChange}
                      label="Семестр"
                      inputProps={{
                        name: 'semester',
                        id: 'outlined-age-native-simple',
                      }}
                    >
                      <MenuItem value="">
                        <em>Любой</em>
                      </MenuItem>
                      <MenuItem value={'1'}>1ый</MenuItem>
                      <MenuItem value={'2'}>2ой</MenuItem>
                      <MenuItem value={'3'}>3ий</MenuItem>
                      <MenuItem value={'4'}>4ый</MenuItem>
                      <MenuItem value={'5'}>5ый</MenuItem>
                      <MenuItem value={'6'}>6ой</MenuItem>
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} sm={12}>
                  <FormControl disabled={semester === ''} size={'small'} required variant="outlined" className={classes.formControl}>
                    <InputLabel htmlFor="subject">Предмет</InputLabel>
                    <Select
                      value={subject}
                      onChange={handleSubjectChange}
                      label="Предмет"
                      inputProps={{
                        name: 'subject',
                        id: 'subject-select',
                      }}
                    >
                      <MenuItem value="">
                        <em>Любой</em>
                      </MenuItem>
                      {props.subjects && props.subjects.map((item) => (
                        <MenuItem value={item.id}>{item.name}</MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                </Grid>
                <Grid item xs={12} sm={6}>
                  <Button
                    href={'#'}
                    fullWidth
                    variant="contained"
                    color="primary"
                  >
                    Создать новый предмет
                  </Button>
                </Grid>
              </Grid>
            </Grid>
          </form>
        </DialogContent>
        <DialogActions>
          <Button onClick={props.handleClose} color="primary">
            Отменить
          </Button>
          <Button onClick={props.handleClose} color="primary">
            Создать
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}
