import React from "react";
import ListItemText from "@material-ui/core/ListItemText";
import Link from "@material-ui/core/Link";
import {Link as RouterLink} from "react-router-dom";

function University(props) {
  return (
    <Link component={RouterLink} to={'/university/' + props.university.id}
          underline="none"
          color="inherit">
      <ListItemText primary={props.university.short}
                    secondary={props.university.full}/>
    </Link>
  );
}

export default University;
