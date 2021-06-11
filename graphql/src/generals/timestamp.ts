import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';

function timestampToDate(t: Timestamp.AsObject): Date {
  let timestamp = new Timestamp();
  timestamp.setSeconds(t.seconds);
  timestamp.setNanos(t.nanos);

  return timestamp.toDate();
}

export { timestampToDate };
