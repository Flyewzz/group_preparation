import React from "react";
import Container from "@material-ui/core/Container";
import {makeStyles} from "@material-ui/core/styles";
import MaterialDetails from "../components/material/MaterialDetails";
import FilesList from "../components/files/FilesList";
import Button from "@material-ui/core/Button";
import SubjectsService from "../services/SubjectsService";
import MaterialService from "../services/MaterialService";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";

const data = {
  name: 'РК №1',
  university: 'МГТУ им. Н. Э. Баумана',
  subject: 'Физика',
  semester: '1ый',
  department: 'ИУ7',
  date: '03.01.2020',
  type: 'РК',
  author: 'username',
  rating: 23,
  description: 'Прекрасное длинное подробное описание восхитительнейших приложенных ниже файлов',
  files: [
    {
      name: 'bilet1.png',
    },
    {
      name: 'bilet2.png',
    },
    {
      name: 'bilet3.png',
    },
    {
      name: 'answers.docx',
    }],
};

const useStyles = makeStyles(() => ({
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
  name: {
    fontWeight: 'bold',
    fontSize: 'xx-large'
  },
  description: {
    fontWeight: 'bold',
    fontSize: 'large',
    marginRight: '5pt',
  },
  descriptionWrapper: {
    margin: '8pt'
  },
  download: {
    marginRight: '5pt',
    width: 'max-content',
    marginLeft: 'auto',
  },
}));

function MaterialPage(props) {
  const styles = useStyles();

  return (
    <Container maxWidth="sm" className={styles.wrapper}>
      <div className={styles.name}>{props.data.name}</div>
      <MaterialDetails material={props.data}/>
      <div className={styles.descriptionWrapper}>
        <span className={styles.description}>Описание:</span>
        <span>{data.description}</span>
      </div>
      <FilesList/>
      <div className={styles.download}>
        <Button variant="contained" color={'primary'}>
          Скачать все
        </Button>
      </div>
    </Container>
  );
}

class MaterialPageController extends React.Component {
  constructor(props) {
    super(props);
    this.materialService = new MaterialService();
  }

  material = {date: ''};

  componentDidMount() {
    this.getMaterial();
  }

  getMaterial = () => {
    const id = this.props.id;
    this.materialService.getById(id).then((result) => {
        runInAction(() => {
          console.log(result);
          this.material = result;
        })
      },
      (error) => {
        console.log(error)
      })
  };

  render() {
    return (
      <MaterialPage data={this.material}/>
    );
  }
}


decorate(MaterialPageController, {
  material: observable,
});

export default observer(MaterialPageController);
