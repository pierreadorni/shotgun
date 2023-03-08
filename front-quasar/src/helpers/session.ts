import { Session } from 'src/types/session';
import axios from 'axios';

const getSession = async (): Promise<Session> => {
  const res = await axios.get('http://localhost:3000/session', { withCredentials: true });
  return res.data;
};

export default getSession;
