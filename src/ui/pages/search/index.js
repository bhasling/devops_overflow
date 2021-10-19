/***************
 * Search Route home page.
 * This page allows a user to search for issues by keywords and view the list of found issues.
 * The user can click on an issue to view the details and the answers of the issue.
 * 
 ***************/
import {React, Component} from 'react';
import Head from 'next/head'
import { connect } from "react-redux";
import IconButton from '@mui/material/IconButton';
import Icon from '@mui/material/Icon';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import NavBar from '../../components/NavBar.js';
import Link from 'next/link';
import RestCall from '../../components/RestCall.js'
import Typography from '@mui/material/Typography';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import SearchSection from '../../components/SearchSection.js';
import DetailSection from '../../components/DetailSection.js';

class SearchPage extends Component {
    constructor(props) {
         super();
         this.props = props;
         this.state = {
             mode: "search",
             issue: undefined
         }
    }

    // Set the current issue in state. This changes the mode from detail to search
    setIssue(issue) {
        if (issue) {
            this.setState({mode: "detail", issue: issue})
        } else {
            this.setState({mode: "search", issue: undefined})
        }
    }
    
    render() {
        return (
            <div>
                <Head>
                    <title>Search Page</title>
                    <meta name="description" content="Searh Page" />
                    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons"/>
                    <link rel="icon" href="/favicon.ico" />
                </Head>
                <NavBar/>
                {this.state.mode == "search"? 
                    (<SearchSection setIssue={(issue) => this.setIssue(issue)}></SearchSection>)
                :
                    (<DetailSection issue={this.state.issue} 
                                    setIssue={(issue) => this.setIssue(issue)}>
                                        Show detail
                    </DetailSection>
                    )
                }
            </div>
        )

    }

}
const mapStateToProps = state => {
    var props = {loginState: state.loginState, searchResult : state.searchResult};
    return props;
}
const mapDispatchToProps = (dispatch) => {
   return {
    updateSearchResult: (searchResult) => dispatch({type: 'updateSearchResult', searchResult:searchResult})
   }
}
export default connect(mapStateToProps, mapDispatchToProps)(SearchPage);


