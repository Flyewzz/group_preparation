import React from 'react';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import CssBaseline from '@material-ui/core/CssBaseline';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import {makeStyles} from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';

const useStyles = makeStyles((theme) => ({
  icon: {
    marginRight: theme.spacing(2),
  },
  cardGrid: {
    paddingTop: theme.spacing(2),
    paddingBottom: theme.spacing(2),
  },
  card: {
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
  },
  cardMedia: {
    paddingTop: '70px',
  },
  cardContent: {
    padding: '5px 5px 0 5px'
  },
  rootName: {
    padding: '0 5px 5px 5px',
  },
  main: {
    width: '100%',
  },
  scroll: {
    height: '345px',
    width: '100%',
    overflowY: 'scroll',
  },
}));

const cards = [1, 2, 3, 4];

function FilesList() {
  const classes = useStyles();

  return (
    <div className={classes.scroll}>
      <CssBaseline/>
      <main className={classes.main}>
        <Container className={classes.cardGrid} maxWidth="sm">
          <Grid container spacing={4}>
            {cards.map((card) => (
              <Grid item key={card} xs={12} sm={6} md={4}>
                <Card className={classes.card}>
                  <CardMedia
                    className={classes.cardMedia}
                    image="https://tritec-education.ru/wp-content/uploads/2015/02/word_thumb.jpg"
                    title="file"
                  />
                  <CardContent className={classes.cardContent}>
                    <Typography classes={{root: classes.rootName}}>
                      Heading
                    </Typography>
                  </CardContent>
                  <CardActions classes={{root: classes.rootName}}>
                    <Button size="small" color="primary">
                      Скачать
                    </Button>
                  </CardActions>
                </Card>
              </Grid>
            ))}
          </Grid>
        </Container>
      </main>
    </div>
  );
}

export default FilesList;