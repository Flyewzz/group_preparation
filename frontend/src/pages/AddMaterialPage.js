import React from "react";
import {observer} from "mobx-react";
import UniversitiesService from "../services/UniversitiesService";
import {decorate, observable, runInAction} from "mobx";
import Container from "@material-ui/core/Container"
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import {makeStyles} from '@material-ui/core/styles';
import InputLabel from "@material-ui/core/InputLabel";
import FormControl from "@material-ui/core/FormControl";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import {DropzoneArea} from 'material-ui-dropzone'
import '../components/material/dropzone.css'
import config from "../config";

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(2),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
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
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  name: {
    fontWeight: 'bold',
    fontSize: 'xx-large'
  },
}));

class DropzoneAreaExample extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      files: []
    };
  }

  handleChange(files) {
    console.log(files);
    this.setState({
      files: files
    });
    this.props.onFileAdd(files);
  }

  render() {
    return (
      <DropzoneArea
        showFileNames
        dropzoneClass={'dropzone'}
        acceptedFiles={['']}
        dropzoneText={'Перетащите сюда необходимые файлы или кликните'}
        onChange={this.handleChange.bind(this)}
      />
    )
  }
}

function AddForm(props) {
  const classes = useStyles();
  const [state, setType] = React.useState('');
  const handleChange = event => {
    setType(event.target.value);
  };

  return (
    <Container maxWidth="sm" className={classes.wrapper}>
      <CssBaseline/>
      <div className={classes.name}>
        Новый материал
      </div>
      <div className={classes.paper}>
        <form onSubmit={props.onSubmit} encType="multipart/form-data" className={classes.form} noValidate>
          <Grid container spacing={2}>
            <Grid item xs={12}>
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
            <Grid item xs={12} sm={9}>
              <FormControl size={'small'} required variant="outlined" className={classes.formControl}>
                <InputLabel htmlFor="university">Университет</InputLabel>
                <Select
                  value={state}
                  onChange={handleChange}
                  label="Университет"
                  inputProps={{
                    name: 'university',
                    id: 'university-select',
                  }}
                >
                  <MenuItem value="">
                    <em>Любой</em>
                  </MenuItem>
                  <MenuItem value={10}>Ten</MenuItem>
                  <MenuItem value={20}>Twenty</MenuItem>
                  <MenuItem value={30}>Thirty</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={3}>
              <FormControl size={'small'} required variant="outlined" className={classes.formControl}>
                <InputLabel htmlFor="outlined-age-native-simple">Семестр</InputLabel>
                <Select
                  value={state}
                  onChange={handleChange}
                  label="Семестр"
                  inputProps={{
                    name: 'subject',
                    id: 'outlined-age-native-simple',
                  }}
                >
                  <MenuItem value="">
                    <em>Любой</em>
                  </MenuItem>
                  <MenuItem value={10}>Ten</MenuItem>
                  <MenuItem value={20}>Twenty</MenuItem>
                  <MenuItem value={30}>Thirty</MenuItem>
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={12} sm={6}>
              <FormControl size={'small'} required variant="outlined" className={classes.formControl}>
                <InputLabel htmlFor="subject">Предмет</InputLabel>
                <Select
                  value={state}
                  onChange={handleChange}
                  label="Предмет"
                  inputProps={{
                    name: 'subject',
                    id: 'subject-select',
                  }}
                >
                  <MenuItem value="">
                    <em>Любой</em>
                  </MenuItem>
                  <MenuItem value={10}>Ten</MenuItem>
                  <MenuItem value={20}>Twenty</MenuItem>
                  <MenuItem value={30}>Thirty</MenuItem>
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
                Создать новый
              </Button>
            </Grid>
            <Grid item xs={12}>
              <TextField
                multiline
                variant="outlined"
                fullWidth
                name="description"
                label="Описание"
                id="description"
              />
            </Grid>
          </Grid>
          <Grid item xs={12}>
            <DropzoneAreaExample onFileAdd={props.onFileAdd}/>
          </Grid>
          <Button
            size={'large'}
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Создать
          </Button>
        </form>
      </div>
    </Container>
  );
}

class AddMaterialPage extends React.Component {
  constructor(props) {
    super(props);

    this.currPage = 1;
    this.pageCount = 1;
    this.universitiesService = new UniversitiesService();
  }

  universities = [];
  files = [];

  componentDidMount() {
    const page = 1; // TODO get from url
    this.getUniversities(page);
  }

  getUniversities = (page) => {
    this.currPage = page;
    this.universitiesService.getPage(page).then((result) => {
        console.log(result);
        runInAction(() => {
          this.pageCount = result.pages;
          this.universities = result.payload;
        });
      },
      (error) => {
        console.log(error);
      });
  };

  onSubmit = async (event) => {
    event.preventDefault();

    const data = new FormData();
    this.files.forEach((file) => {
      data.append('file', file);
      data.append('filename', file.name);
    });
  };

  onFileAdd = (files) => {
    this.files = files;
  };

  render() {
    return (
      <AddForm onSubmit={this.onSubmit}
               onFileAdd={this.onFileAdd}/>
    );
  }
}

decorate(AddMaterialPage, {
  pageCount: observable,
  universities: observable,
});

export default observer(AddMaterialPage);
