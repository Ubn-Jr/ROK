import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View, Image, Button } from 'react-native';

export default function App() {
  return (
    <View style={styles.container}>
      <Text></Text>
      <StatusBar style="auto" />
  
      <Image
        
        source={{
          uri: 'https://www.doppelme.com/PM11101:C1471:A2FFDAB9/avatar.png',
          width:150,
          height:300,
        }}
      />
<Image 
        source={{
          uri: 'https://www.doppelme.com/PM11102:C1471:A2FFDAB9/avatar.png',
          width:150,
          height:300,
        }}
      
      />
       <View style={styles.buttonContainer}>
       <Button title="ROK" onPress={() => console.log('Butona tıklandı!')} />
      </View>
     
    </View>
     

  
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    justifyContent: 'center', 
    alignItems: 'center', 
    padding: 10,
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
  buttonContainer: {
    position: 'absolute',
    bottom: 200,
    left: 0,
    right: 0,
    justifyContent: 'center',
    alignItems: 'center',
  },
});
