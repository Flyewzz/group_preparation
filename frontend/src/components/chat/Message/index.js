import React from 'react';
import moment from 'moment';
import './Message.css';
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles(() => ({
  author: {
    color: '#007aff'
  },
}));

export default function Message(props) {
  const classes = useStyles(props);
    const {
      data,
      isMine,
      author,
      startsSequence,
      endsSequence,
      showTimestamp
    } = props;

    const friendlyTimestamp = moment(data.timestamp).format('LLLL');
    return (
      <div className={[
        'message',
        `${isMine ? 'mine' : ''}`,
        `${startsSequence ? 'start' : ''}`,
        `${endsSequence ? 'end' : ''}`
      ].join(' ')}>
        {
          showTimestamp &&
            <div className="timestamp">
              { friendlyTimestamp }
            </div>
        }

        <div className="bubble-container">
          <div className="bubble" title={friendlyTimestamp}>
            <div className={classes.author}>{!isMine && startsSequence && author + ':'}</div>
            { data.message }
          </div>
        </div>
      </div>
    );
}