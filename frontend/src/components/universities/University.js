import React from "react";
import ListItemText from "@material-ui/core/ListItemText";
import ListItemIcon from "@material-ui/core/ListItemAvatar"
import Link from "@material-ui/core/Link";
import {Link as RouterLink} from "react-router-dom";
import Avatar from "@material-ui/core/Avatar";
import config from "../../config.js"
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
  link: {
    display: 'flex',
  },
  avatar: {
    width: theme.spacing(14),
    height: theme.spacing(14),
  },
  primaryText: {
    marginTop: '15pt',
    fontSize: 'x-large',
  },
  secondaryText: {
    fontSize: 'large',
  },
}));

function University(props) {
  const classes = useStyles();
  return (
    <Link component={RouterLink} to={'/university/' + props.university.id}
          underline="none"
          className={classes.link}
          color="inherit">
      <ListItemIcon>
        <Avatar
          alt={`icon`}
          variant={'square'}
          src={`${config.apiUrl}/university/${props.university.id}/avatar`}
          className={classes.avatar}
        />
      </ListItemIcon>
      <ListItemText primary={props.university.name}
                    primaryTypographyProps={{className: classes.primaryText}}
                    secondary={props.university.full_name}
                    secondaryTypographyProps={{className: classes.secondaryText}}/>
    </Link>
  );
}

export default University;
