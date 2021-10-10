import '../styles/globals.css'
import store from '../components/store';
import { Provider } from 'react-redux';
import { createTheme, ThemeProvider } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    primary: {
      // Purple and green play nicely together.
      main: '#025608'
    },
    secondary: {
      // This is green.A700 as hex.
      main: '#11cb5f',
    },
  },
});

function MyApp({ Component, pageProps }) {
  return <Provider store={store}><ThemeProvider theme={theme}><Component {...pageProps} /></ThemeProvider></Provider>
}

export default MyApp
