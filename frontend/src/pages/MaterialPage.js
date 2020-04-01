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
import SearchInput from "../components/common/SearchInput";
import MaterialPathHeader from "../components/material/MaterialPathHeader";
import InputBase from "@material-ui/core/InputBase";
import {Search} from "@material-ui/icons";

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
    display: 'flex',
    flexDirection: 'row',
    height: '100%',
  },
  left: {
    marginLeft: '55pt',
  },
  right: {
    marginLeft: '30pt',
    marginTop: '10pt',
    width: '100%',
    height: '100%',
  },
  description: {
    fontWeight: 'bold',
    fontSize: 'large',
    marginRight: '5pt',
  },
  descriptionWrapper: {
    margin: '8pt'
  },
  filesHead: {
    display: 'flex',
    flexDirection: 'row',
  },
  inputBox: {
    display: 'flex',
    alignItems: 'center',
    border: '1px solid lightgray',
    borderRadius: '4px',
    padding: '3pt 0 0 2pt',
    marginRight: '15pt',
    width: '40%'
  },
  input: {
    paddingLeft: '2pt',
  },
}));

function MaterialPage(props) {
  const styles = useStyles();

  return (
    <>
      <MaterialPathHeader university={props.university}
                          subject={props.subject}
                          materialName={props.data.name}/>
      <main className={styles.wrapper}>
        <div className={styles.left}>
          <MaterialDetails material={props.data}/>
          <div className={styles.descriptionWrapper}>
            <span className={styles.description}>Описание:</span>
            <span>{data.description}</span>
          </div>
        </div>
        <div className={styles.right}>
          <div className={styles.filesHead}>
            <div className={styles.inputBox}>
              <Search/>
              <InputBase fullWidth
                         onChange={props.onChange}
                         placeholder={props.placeholder}
                         className={styles.input}
                         inputProps={{'aria-label': 'search'}}
              />
            </div>
            <Button variant="contained" color={'primary'}>
              Скачать все
            </Button>
          </div>
          <FilesList/>
        </div>
      </main>
    </>
  );
}

class MaterialPageController extends React.Component {
  constructor(props) {
    super(props);
    this.materialService = new MaterialService();
  }

  university = {id: -1, name: ''};
  subject = {id: -1, name: ''};
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
      <MaterialPage university={this.university}
                    subject={this.subject}
                    data={this.material}/>
    );
  }
}


decorate(MaterialPageController, {
  material: observable,
});

export default observer(MaterialPageController);
