/***************
 * Top level home page route
 * 
 * Home page for the top level of the application.
 ***************/
import Head from 'next/head'
import Typography from '@mui/material/Typography';
import NavBar from '../components/NavBar.js';
import LinkCard from '../components/LinkCard.js';

export default function Home() {
  const cardContainerStyle = {
      'fontSize':'20px',
      'display': 'flex',
      'flexDirection' : 'row',
      'justifyContent': 'space-between',
      'alignContent': 'stretch',
      'columnGap': '30px',
      'cursor':'pointer'
  }

  const circleIconStyle = {
      'borderRadius':'50%',
      'backgroundColor': '#FDA9F7',
      'filter': 'invert(100%)'
  }

  const containerStyle = {
    'minHeight':'100vh',
    'padding': '0 0.5rem',
    'display' : 'flex',
    'flexDirection': 'column',
    'justisfyContent': 'center',
    'alignItems': 'center',
    'height':'100vh'
}

  const mainStyle = {
    'flex':'1',
    'display' : 'flex',
    'flexDirection': 'column',
    'justisfyContent': 'top',
    'alignItems': 'center'
  }

  const footerStyle = {
    'width': '100%',
    'height': '60px',
    'borderTop':'1px solid #eaeaea',
    'display' : 'flex',
    'justisfyContent': 'center',
    'alignItems': 'center'
  }


  return (

    <div style={containerStyle}>
      <Head>
        <title>Home Page</title>
        <meta name="description" content="Home Page" />
        <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons"/>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <NavBar></NavBar>
      <div style={mainStyle}>
          <Typography color='primary' fontWeight = 'bold' fontSize = '2rem'>
          <div >Dev Ops Overflow</div>
          </Typography>
          <div>
              <img  style={circleIconStyle} src='/infinity_icon.png' height='50px'></img>
          </div>
          <div>
            <div style={cardContainerStyle}>
              <LinkCard 
                name={'Search'}
                title={'Search for questions or issues previous.'}
                detail={
                  `
                  The search feature will
                    search for any issue or solution that contains the search words. You can
                    browse issues found and view all the solutions.
                    While browsing you can link to similar issues.
                    You can also submit new solutions to the issue or mark previous solutions as in-accurate.
                  `
                }
                route={'/search'}                  
                />

              <LinkCard 
                name={'Submit'}
                title={'Submit a new problem or issue.'}
                detail={
                  `
                  You can submit the text of a new issue and associate that issue with a product and version.
                  `
                }
                route={'/submit'}
                />
            </div>
          </div>
      </div>

      <footer style={footerStyle}>
        <Typography  fontSize='0.7rem' color='primary'>
        Dev Ops Overflow is tool for submiting and searching for issues and resolutions of devops issues for local products to support DevOps.
        </Typography>
      </footer>
    </div>
  )
}
