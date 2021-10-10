/****************
 * SignInButton.js
 * 
 */
 import {React, Component} from 'react';
 import { connect } from "react-redux";
 
 class SignInButton extends Component {
     constructor(props) {
         super();
         this.props = props;
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
        if (this.props.loginState.hover) {
            nameStyle['textDecoration'] = 'underline';
        }
        return(
            <div style={{'display':'inline-block', 'textAlign': 'left'}} 
                onClick={()=>this.clickPeople()}
                onMouseEnter={()=>this.mouseEnter()}
                onMouseLeave={()=>this.mouseLeave()}
            >
                <div>
                <span >
                    <span style={iconStyle}>{this.props.loginState.hover ? 'person_outline' : 'person'}</span>
                    <span style={{'display':'inline-block'}}>
                        <div style={nameStyle} >{this.props.loginState.loginId? this.props.loginState.loginUser: 'My Account'}</div>
                        <div style={{'cursor':'pointer'}} >
                            {this.props.loginState.loginId? 'Sign Out': 'Sign In'}
                        </div>
                    </span>
                </span>
                </div>
            </div>
        );
     }

     mouseEnter() {
        this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'hover': true});

     }
     mouseLeave() {
        this.props.updateLoginState({'loginId': this.props.loginState.loginId, 'loginUser': this.props.loginState.loginUser, 'hover': false});
    }
    clickPeople() {
        if (this.props.loginState.loginId) {
            this.props.updateLoginState({'loginId': null, 'loginUser': null, 'hover': false});
        } else {
            this.props.updateLoginState({'loginId': '2', 'loginUser': 'myusername', 'hover': false});
        }
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
