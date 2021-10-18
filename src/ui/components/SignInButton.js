/****************
 * SignInButton.js
 * 
 * Displays the name of the logged in user or a SignIn button to login.
 * The SignIn button shows a login dialog where the user can login. This
 * post a request to the server with a user/password and returns with
 * the name of the user and an authentication cookie to allow other REST
 * requests to work.
 * 
 */
import {React, Component} from 'react';
import { connect } from "react-redux";
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
 
class SignInButton extends Component {
     constructor(props) {
         super();
         this.props = props;
         this.state = {
             loginOpen: false,
             signUpOpen: false
         }
     }
 
 
     render() {
        var iconStyle = {
            'fontSize':'20px',
            'display': 'inline-block',
            'width': '20px',
            'marginTop' : '4px',
            'verticalAlign':'top',
            'fontFamily': 'Material Icons',
            'cursor':'pointer'
        }
        var nameStyle = {
            'fontSize':'16px',
            'marginTop' : '4px',
            'verticalAlign':'middle',
            'display': 'inline-block',
            'fontWeight': 'bold',
            'cursor':'pointer'
        }
        var signUpoutStyle = {
            'cursor':'pointer'
        }
        if (this.props.loginState.userHover) {
            nameStyle['textDecoration'] = 'underline';
        }
        if (this.props.loginState.signUpOutHover) {
            signUpoutStyle['textDecoration'] = 'underline';
        }
        return(
            <div style={{'display':'inline-block', 'textAlign': 'left'}}                
            >
                <div>
                    <span>
                        <span style={iconStyle}>{this.props.loginState.userHover ? 'person_outline' : 'person'}</span>
                        <span style={{'display':'inline-block'}}>
                            <div style={nameStyle} onClick={()=>this.clickPeople()} onMouseEnter={()=>this.userMouseEnter()} onMouseLeave={()=>this.userMouseLeave()}>
                                {this.props.loginState.loginId? this.props.loginState.loginUser: 'Sign In'}
                            </div>
                            <div style={signUpoutStyle}  onClick={()=>this.clickSignUpOut()} onMouseEnter={()=>this.signOutUpMouseEnter()} onMouseLeave={()=>this.signOutUpMouseLeave()}>
                                {this.props.loginState.loginId? 'Sign Out': 'Sign Up'}
                            </div>
                        </span>
                    </span>

                    <Dialog open={this.state.loginOpen} onClose={()=>this.closeLoginDialog()}>
                        <DialogTitle>Login</DialogTitle>
                        <DialogContent>
                            <DialogContentText>
                                To login, please enter your user name nad password.
                            </DialogContentText>
                            <TextField
                                autoFocus
                                margin="dense"
                                id="name"
                                label="User Id"
                                type="string"
                                fullWidth
                                variant="standard"
                            />
                            <TextField
                                margin="dense"
                                id="password"
                                label="Password"
                                type="password"
                                fullWidth
                                variant="standard"
                            />
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={()=>this.closeLoginDialog()}>Cancel</Button>
                            <Button onClick={()=>this.closeLoginDialog()}>Login</Button>
                        </DialogActions>
                    </Dialog>                
                    <Dialog open={this.state.signUpOpen} onClose={()=>this.closeSignupDialog()}>
                        <DialogTitle>Sign Up</DialogTitle>
                        <DialogContent>
                            <DialogContentText>
                                Enter your desired user and password to sign up.
                            </DialogContentText>
                            <TextField
                                autoFocus
                                margin="dense"
                                id="name"
                                label="User Id"
                                type="string"
                                fullWidth
                                variant="standard"
                            />
                            <TextField
                                margin="dense"
                                id="password"
                                label="Password"
                                type="password"
                                fullWidth
                                variant="standard"
                            />
                        </DialogContent>
                        <DialogActions>
                            <Button onClick={()=>this.closeSignupDialog()}>Cancel</Button>
                            <Button onClick={()=>this.closeSignupDialog()}>Sign Up</Button>
                        </DialogActions>
                    </Dialog>                

                </div>
            </div>
        );
     }

     userMouseEnter() {
        if (!this.props.loginState.loginId) {
            this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'userHover': true});
        }
     }
     userMouseLeave() {
        if (!this.props.loginState.loginId) {
            this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'userHover': false});
        }
    }
    signOutUpMouseEnter() {
        this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'signUpOutHover': true});
     }
     signOutUpMouseLeave() {
        this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'signUpOutHover': false});
    }
    clickPeople() {
        if (this.props.loginState.loginId) {
            this.props.updateLoginState({'loginId': null, 'loginUser': null, 'userHover': false});
        } else {
            this.showLoginDialog();
        }
    }
    showLoginDialog() {
        this.setState({loginOpen: true});
    }
    closeLoginDialog() {
        if (this.props.loginState.loginId) {
            this.setState({loginOpen: false});
            this.props.updateLoginState({'loginId': null, 'loginUser': null, 'userHover': false});
        } else {
            this.setState({loginOpen: false});
            this.props.updateLoginState({'loginId': '2', 'loginUser': 'myusername', 'userHover': false});
        }
    }
    clickSignUpOut() {
        if (this.props.loginState.loginId) {
            this.props.updateLoginState({'loginId': null, 'loginUser': null, 'userHover': false});
        } else {
            this.showSignupDialog();
        }
    }
    showSignupDialog() {
        this.setState({signUpOpen: true});
    }
    closeSignupDialog() {
        if (this.props.loginState.loginId) {
        } else {
        }
        this.setState({signUpOpen: false});
    }

}
 const mapStateToProps = state => {
    var props = {loginState: state.loginState};
    return props;
}
const mapDispatchToProps = (dispatch) => {
   return {
     updateLoginState: (loginState) => dispatch({type: 'updateLoginState', loginState:loginState}),
     setGeneralMessage: (message) => dispatch({type: 'setGeneralMessage', general_message:message})
   }
}
export default connect(mapStateToProps, mapDispatchToProps)(SignInButton);
